/*
	Copyright 2023 Loophole Labs

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

		   http://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

package main

import (
	"context"
	"flag"
	"github.com/extism/extism"
	"github.com/loopholelabs/polyglot-go"
	"github.com/loopholelabs/scale-benchmarks/http/compile"
	native "github.com/loopholelabs/scale-benchmarks/pkg/native/go"
	"github.com/loopholelabs/scale-benchmarks/pkg/scale/go/signature/text-signature"
	scale "github.com/loopholelabs/scale/go"
	"github.com/loopholelabs/scale/go/tests/harness"
	"github.com/loopholelabs/scalefile"
	"github.com/loopholelabs/scalefile/scalefunc"
	"github.com/valyala/fasthttp"
	"log"
	"os"
	"os/exec"
	"sync"
	"unsafe"
)

var (
	kind string
	addr string
)

func main() {
	flag.StringVar(&kind, "kind", "native", "The kind of benchmark to run (native, scale, extism)")
	flag.StringVar(&addr, "addr", ":8080", "The listen address for the server")

	flag.Parse()

	if addr == "" {
		panic("addr is required")
	}

	srv := &fasthttp.Server{
		Handler:               nil,
		ReadBufferSize:        65536 * 2,
		WriteBufferSize:       65536 * 2,
		TCPKeepalive:          true,
		NoDefaultServerHeader: true,
		NoDefaultDate:         true,
		NoDefaultContentType:  true,
		StreamRequestBody:     true,
	}

	switch kind {
	case "native":
		srv.Handler = func(ctx *fasthttp.RequestCtx) {
			ctx.SetContentType("text/plain")
			b := ctx.Request.Body()
			if len(b) > 0 {
				r, err := native.ReplaceAll(unsafe.String(&b[0], len(b)))
				if err != nil {
					ctx.SetStatusCode(fasthttp.StatusInternalServerError)
					_, _ = ctx.WriteString(err.Error())
					return
				}
				_, _ = ctx.WriteString(r)
			} else {
				_, _ = ctx.WriteString("Hello, World!")
			}
		}
	case "scale":
		moduleConfig := &harness.Module{
			Name:          "text_signature",
			Path:          "../pkg/scale/rust/modules/text_signature/text_signature.rs",
			Signature:     "text_signature",
			SignaturePath: "../../../signature/text-signature",
		}

		generatedModules, err := compile.RustSetup(
			[]*harness.Module{moduleConfig},
			[]*scalefile.Dependency{
				{
					Name:    "scale_signature",
					Version: "0.2.1",
				},
				{
					Name:    "regex",
					Version: "1.7.1",
				},
				{
					Name:    "lazy_static",
					Version: "1.4.0",
				},
			},
		)
		if err != nil {
			panic(err)
		}

		module, err := os.ReadFile(generatedModules[moduleConfig])
		if err != nil {
			panic(err)
		}

		scaleFunc := &scalefunc.ScaleFunc{
			Version:   scalefunc.V1Alpha,
			Name:      "TestName",
			Tag:       "TestTag",
			Signature: "ExampleName@ExampleVersion",
			Language:  scalefunc.Rust,
			Function:  module,
		}

		r, err := scale.NewWithSignature(context.Background(), text.New, []*scalefunc.ScaleFunc{scaleFunc})
		if err != nil {
			panic(err)
		}

		srv.Handler = func(ctx *fasthttp.RequestCtx) {
			ctx.SetContentType("text/plain")
			b := ctx.Request.Body()
			if len(b) > 0 {
				i, err := r.Instance(nil)
				if err != nil {
					ctx.SetStatusCode(fasthttp.StatusInternalServerError)
					_, _ = ctx.WriteString(err.Error())
					return
				}

				i.Context().Data = unsafe.String(&b[0], len(b))

				err = i.Run(ctx)
				if err != nil {
					ctx.SetStatusCode(fasthttp.StatusInternalServerError)
					_, _ = ctx.WriteString(err.Error())
					return
				}
				_, _ = ctx.WriteString(i.Context().Data)
			} else {
				_, _ = ctx.WriteString("Hello, World!")
			}
		}
	case "extism":
		cmd := exec.Command("cargo", "build", "--release", "--target", "wasm32-unknown-unknown")

		cmd.Dir = "../pkg/extism/rust"

		if err := cmd.Run(); err != nil {
			panic(err)
		}

		ctx := extism.NewContext()
		defer ctx.Free()

		manifest := extism.Manifest{Wasm: []extism.Wasm{extism.WasmFile{Path: "../pkg/extism/rust/target/wasm32-unknown-unknown/release/rust.wasm"}}}

		pool := &sync.Pool{
			New: func() interface{} {
				plugin, err := ctx.PluginFromManifest(manifest, []extism.Function{}, true)
				if err != nil {
					panic(err)
				}
				return plugin
			},
		}

		srv.Handler = func(ctx *fasthttp.RequestCtx) {
			ctx.SetContentType("text/plain")
			b := ctx.Request.Body()
			if len(b) > 0 {
				buf := polyglot.GetBuffer()
				polyglot.Encoder(buf).String(unsafe.String(&b[0], len(b)))
				plugin, ok := pool.Get().(extism.Plugin)
				if !ok {
					ctx.SetStatusCode(fasthttp.StatusInternalServerError)
					_, _ = ctx.WriteString("could not get plugin from pool")
					return
				}
				out, err := plugin.Call("match_regex", buf.Bytes())
				polyglot.PutBuffer(buf)
				if err != nil {
					ctx.SetStatusCode(fasthttp.StatusInternalServerError)
					_, _ = ctx.WriteString(err.Error())
					return
				}

				dec := polyglot.GetDecoder(out)
				matches, err := dec.String()
				if err != nil {
					ctx.SetStatusCode(fasthttp.StatusInternalServerError)
					_, _ = ctx.WriteString(err.Error())
					return
				}
				dec.Return()
				pool.Put(plugin)
				_, _ = ctx.WriteString(matches)
			} else {
				_, _ = ctx.WriteString("Hello, World!")
			}
		}
	default:
		panic("invalid kind")
	}

	log.Printf("Listening on %s", addr)
	if err := srv.ListenAndServe(addr); err != nil {
		panic(err)
	}
}
