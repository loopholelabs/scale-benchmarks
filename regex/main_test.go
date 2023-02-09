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
	"github.com/extism/extism"
	"github.com/loopholelabs/polyglot-go"
	native "github.com/loopholelabs/scale-benchmarks/pkg/native/go"
	"github.com/loopholelabs/scale-benchmarks/pkg/scale/go/signature/text-signature"
	"github.com/loopholelabs/scale-benchmarks/regex/compile"
	scale "github.com/loopholelabs/scale/go"
	"github.com/loopholelabs/scale/go/tests/harness"
	"github.com/loopholelabs/scalefile"
	"github.com/loopholelabs/scalefile/scalefunc"
	"os"
	"os/exec"
	"testing"
)

const (
	Regex2048 = "BYwnOjWhprPmDncp8qpQ5CY4r1RGZuqKLBowmtMCd test ETjLOG685YC4RIjXB0HadNpqYS4M7GPGUVAKRZRC1ibqQqGnuzqX2Hjosm6MKNCp5QifX7Up2phqkFqkjpSu3k59oi6M5YbTMiy4JukVFx2402IlrHU1McK7US0skB1cF0W2ZDpsypNmGJRXRMY0pPsYbw7G2a0xJnhTITXcuF5xJWR1rz5zdGZQbbjZoHZcEnveDFq5kOmCVc test DsJVHTsAlypLI9sVtbTLwmE1DG2C6AgUo3GO1DpCx3jV43oXUxaTVJqZO13AYqvNPbxizYZ5BckZFBbJybY3Vnm20Sm7nXbwZs5N2ugz3EpUQvXwqHdHWzc1T8uKPD5LTDM8UBpVoF test 9G3mWarrp43SvoidITriFhzHmyVWNd6n2LIVocr3pOai4DOlkAn7QDup6z6spMAf8UcI4wbfoSzG0k5Qy1rGBhPaJKJRW2 MC9ma3U3rnjAOBtEUHZ2qfOUpfMNgPlGpvzr4IGNNFf9RFlF7yRUBvRnYxyonIWPPiR1x1wWgxc20o5cW4GU7kytAOuGlpzpykcAxCJLLP6wJegaMhAeb8xBLpuBetNEbfcyyOcJBun5BhmFOmv8 test IvICWx2wlYZ61YDBpPcIpqnMb9MHwT8GroC1YITZBlNGBHMpAe4d2sNZe9d0Wvfbv5mMo30Bm1Pa5S3x38jgu6y0BaqZl9GhlukE9CqPJGUsJZ5suDH19WiOrvz7mXwXhi4lWm1YdwNi0xhVnXITtmKq5rikIS6dul1USgDf3TwyLYpyCG46Xj92PssJmnhPdH1WAnvXY sbs8RaemyqmPggtGNwU2JjuPjdmQRakIusv2WimN7zG8R8Pf1225IAJ2j8aiZBrxnjmrucaYOQCrLm7e2Q5q8 test HOkCEJJGHVLYJtGgHKa1PRQ5qCcsIAUdkW3yRfdulutteLe3We9z9XQvWuTYMLDPpOJqMzDNTGpTYts7AL8pFog1k82XVuMZ6ItccxOBpuzDcahH4wDqCGjak8qPVxmnrGmSsrdUHVz6SrScElMo0nOF8RIpYAVdJr5NxWIK1uzc1iIiZnbUD6uDNmBkmfec6IgK6aqnEZaGLDJXDHSYfzWUOi7y3KNPl0CghL9BId8v4040mCKMfmdthWWLJ2tpWIo1482ghiU5 2qtrzgFgYKfyfr4X6FXzN3hM3bLnuwItQrTCEp3BYz79bCAaQGhicZzqE83Mh2 test IIVID622qlEyVEGuEmNJ5JteEzbpklhTKnVMflzzWyWbZe6kIgeUr9mxWjkJGisvRbZKwfnojeC82M1nHgUa4k46x7Dw7mL3rChORjBxBMYjFeOvCsT6kEo3vPeachLUKdkExJbr9Yei0fKyOFSDlxpFhlRKuwGxXu4jGo4CzKDsVsahqzC9iGw53bHiw0V4Pwmdhzv482s3zU9XLTgQr6GuL1I0kSfh9BkVoK5fFvg1hm7ECrt6p8q3kLVjxte EK9W9q2q9etMaPLymcCRZ0XauMDzJY08JeVvovnT2g5hxE7UGW1 test YRotQUivrrXQnhEw55faznZZBU1ULVs4BfYkIkEfS91NetBhona6zrzDwMsXi0FJjdaiJ25lvetPDaMzUs0l6nfkGkVyU376mFPfPkpBKZR2z2Xwzxndi0SkUnqm8jCa7iq2oSJstTdUXtCK2xTXMIh7tiuPVftit GFYQXXI3vY QFe1xShWJgFAqYguQ8gcxMPSzMlyDaPmMuTPgFZDM0cd test NS3fTggxBa4p5jgS4S0nhae05RkYkXGzuNMXeu6IoR9PFqVFnXcBYD0Ld9otrAiqUuIGYGmAjm3Wx29va2UtIFaRhL02ckRfycz3BGfwqYl3TGtjWdKjmxn1WreRIIq5gkbWJws5VQsov0V2U8pGedj N2RDqWgh2tFiJA9fmytgRgqSnqxIwyBMgY5RnE6CZ0 test Iv4QPiWMu0oG70e4nSNtG13O test "
	Match2048 = "BYwnOjWhprPmDncp8qpQ5CY4r1RGZuqKLBowmtMCd wasm ETjLOG685YC4RIjXB0HadNpqYS4M7GPGUVAKRZRC1ibqQqGnuzqX2Hjosm6MKNCp5QifX7Up2phqkFqkjpSu3k59oi6M5YbTMiy4JukVFx2402IlrHU1McK7US0skB1cF0W2ZDpsypNmGJRXRMY0pPsYbw7G2a0xJnhTITXcuF5xJWR1rz5zdGZQbbjZoHZcEnveDFq5kOmCVc wasm DsJVHTsAlypLI9sVtbTLwmE1DG2C6AgUo3GO1DpCx3jV43oXUxaTVJqZO13AYqvNPbxizYZ5BckZFBbJybY3Vnm20Sm7nXbwZs5N2ugz3EpUQvXwqHdHWzc1T8uKPD5LTDM8UBpVoF wasm 9G3mWarrp43SvoidITriFhzHmyVWNd6n2LIVocr3pOai4DOlkAn7QDup6z6spMAf8UcI4wbfoSzG0k5Qy1rGBhPaJKJRW2 MC9ma3U3rnjAOBtEUHZ2qfOUpfMNgPlGpvzr4IGNNFf9RFlF7yRUBvRnYxyonIWPPiR1x1wWgxc20o5cW4GU7kytAOuGlpzpykcAxCJLLP6wJegaMhAeb8xBLpuBetNEbfcyyOcJBun5BhmFOmv8 wasm IvICWx2wlYZ61YDBpPcIpqnMb9MHwT8GroC1YITZBlNGBHMpAe4d2sNZe9d0Wvfbv5mMo30Bm1Pa5S3x38jgu6y0BaqZl9GhlukE9CqPJGUsJZ5suDH19WiOrvz7mXwXhi4lWm1YdwNi0xhVnXITtmKq5rikIS6dul1USgDf3TwyLYpyCG46Xj92PssJmnhPdH1WAnvXY sbs8RaemyqmPggtGNwU2JjuPjdmQRakIusv2WimN7zG8R8Pf1225IAJ2j8aiZBrxnjmrucaYOQCrLm7e2Q5q8 wasm HOkCEJJGHVLYJtGgHKa1PRQ5qCcsIAUdkW3yRfdulutteLe3We9z9XQvWuTYMLDPpOJqMzDNTGpTYts7AL8pFog1k82XVuMZ6ItccxOBpuzDcahH4wDqCGjak8qPVxmnrGmSsrdUHVz6SrScElMo0nOF8RIpYAVdJr5NxWIK1uzc1iIiZnbUD6uDNmBkmfec6IgK6aqnEZaGLDJXDHSYfzWUOi7y3KNPl0CghL9BId8v4040mCKMfmdthWWLJ2tpWIo1482ghiU5 2qtrzgFgYKfyfr4X6FXzN3hM3bLnuwItQrTCEp3BYz79bCAaQGhicZzqE83Mh2 wasm IIVID622qlEyVEGuEmNJ5JteEzbpklhTKnVMflzzWyWbZe6kIgeUr9mxWjkJGisvRbZKwfnojeC82M1nHgUa4k46x7Dw7mL3rChORjBxBMYjFeOvCsT6kEo3vPeachLUKdkExJbr9Yei0fKyOFSDlxpFhlRKuwGxXu4jGo4CzKDsVsahqzC9iGw53bHiw0V4Pwmdhzv482s3zU9XLTgQr6GuL1I0kSfh9BkVoK5fFvg1hm7ECrt6p8q3kLVjxte EK9W9q2q9etMaPLymcCRZ0XauMDzJY08JeVvovnT2g5hxE7UGW1 wasm YRotQUivrrXQnhEw55faznZZBU1ULVs4BfYkIkEfS91NetBhona6zrzDwMsXi0FJjdaiJ25lvetPDaMzUs0l6nfkGkVyU376mFPfPkpBKZR2z2Xwzxndi0SkUnqm8jCa7iq2oSJstTdUXtCK2xTXMIh7tiuPVftit GFYQXXI3vY QFe1xShWJgFAqYguQ8gcxMPSzMlyDaPmMuTPgFZDM0cd wasm NS3fTggxBa4p5jgS4S0nhae05RkYkXGzuNMXeu6IoR9PFqVFnXcBYD0Ld9otrAiqUuIGYGmAjm3Wx29va2UtIFaRhL02ckRfycz3BGfwqYl3TGtjWdKjmxn1WreRIIq5gkbWJws5VQsov0V2U8pGedj N2RDqWgh2tFiJA9fmytgRgqSnqxIwyBMgY5RnE6CZ0 wasm Iv4QPiWMu0oG70e4nSNtG13O wasm "

	Regex4096 = Regex2048 + Regex2048
	Match4096 = Match2048 + Match2048

	Regex8192 = Regex4096 + Regex4096
	Match8192 = Match4096 + Match4096

	Regex16384 = Regex8192 + Regex8192
	Match16384 = Match8192 + Match8192

	Regex32768 = Regex16384 + Regex16384
	Match32768 = Match16384 + Match16384

	Regex65536 = Regex32768 + Regex32768
	Match65536 = Match32768 + Match32768
)

func BenchmarkScaleRust(b *testing.B) {
	moduleConfig := &harness.Module{
		Name:          "text_signature",
		Path:          "../pkg/scale/rust/modules/text_signature/text_signature.rs",
		Signature:     "text_signature",
		SignaturePath: "../../../signature/text-signature",
	}

	generatedModules := compile.RustSetup(
		b,
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

	ctx := context.Background()
	b.ResetTimer()
	b.Run("match regex 2048", func(b *testing.B) {
		b.SetBytes(int64(len(Regex2048)))
		b.ReportAllocs()
		for j := 0; j < b.N; j++ {
			i, err := r.Instance(nil)
			if err != nil {
				panic(err)
			}

			i.Context().Data = Regex2048

			if err := i.Run(ctx); err != nil {
				panic(err)
			}

			if i.Context().Data != Match2048 {
				panic("invalid regex match")
			}
		}
	})

	b.Run("match regex 4096", func(b *testing.B) {
		b.SetBytes(int64(len(Regex4096)))
		b.ReportAllocs()
		for j := 0; j < b.N; j++ {
			i, err := r.Instance(nil)
			if err != nil {
				panic(err)
			}

			i.Context().Data = Regex4096

			if err := i.Run(ctx); err != nil {
				panic(err)
			}

			if i.Context().Data != Match4096 {
				panic("invalid regex match")
			}
		}
	})

	b.Run("match regex 8192", func(b *testing.B) {
		b.SetBytes(int64(len(Regex8192)))
		b.ReportAllocs()
		for j := 0; j < b.N; j++ {
			i, err := r.Instance(nil)
			if err != nil {
				panic(err)
			}

			i.Context().Data = Regex8192

			if err := i.Run(ctx); err != nil {
				panic(err)
			}

			if i.Context().Data != Match8192 {
				panic("invalid regex match")
			}
		}
	})

	b.Run("match regex 16384", func(b *testing.B) {
		b.SetBytes(int64(len(Regex16384)))
		b.ReportAllocs()
		for j := 0; j < b.N; j++ {
			i, err := r.Instance(nil)
			if err != nil {
				panic(err)
			}

			i.Context().Data = Regex16384

			if err := i.Run(ctx); err != nil {
				panic(err)
			}

			if i.Context().Data != Match16384 {
				panic("invalid regex match")
			}
		}
	})

	b.Run("match regex 32768", func(b *testing.B) {
		b.SetBytes(int64(len(Regex32768)))
		b.ReportAllocs()
		for j := 0; j < b.N; j++ {
			i, err := r.Instance(nil)
			if err != nil {
				panic(err)
			}

			i.Context().Data = Regex32768

			if err := i.Run(ctx); err != nil {
				panic(err)
			}

			if i.Context().Data != Match32768 {
				panic("invalid regex match")
			}
		}
	})

	b.Run("match regex 65536", func(b *testing.B) {
		b.SetBytes(int64(len(Regex65536)))
		b.ReportAllocs()
		for j := 0; j < b.N; j++ {
			i, err := r.Instance(nil)
			if err != nil {
				panic(err)
			}

			i.Context().Data = Regex65536

			if err := i.Run(ctx); err != nil {
				panic(err)
			}

			if i.Context().Data != Match65536 {
				panic("invalid regex match")
			}
		}
	})
}

func BenchmarkExtismRust(b *testing.B) {
	cmd := exec.Command("cargo", "build", "--release", "--target", "wasm32-unknown-unknown")

	cmd.Dir = "../pkg/extism/rust"

	if err := cmd.Run(); err != nil {
		panic(err)
	}

	ctx := extism.NewContext()
	defer ctx.Free()

	manifest := extism.Manifest{Wasm: []extism.Wasm{extism.WasmFile{Path: "../pkg/extism/rust/target/wasm32-unknown-unknown/release/rust.wasm"}}}

	plugin, err := ctx.PluginFromManifest(manifest, []extism.Function{}, true)
	if err != nil {
		panic(err)
	}

	buf := polyglot.NewBuffer()
	b.ResetTimer()
	b.Run("match regex 2048", func(b *testing.B) {
		polyglot.Encoder(buf).String(Regex2048)
		b.SetBytes(int64(len(Regex2048)))
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			out, err := plugin.Call("match_regex", buf.Bytes())
			if err != nil {
				panic(err)
			}

			dec := polyglot.GetDecoder(out)
			matches, err := dec.String()
			if err != nil {
				panic(err)
			}
			dec.Return()

			if matches != Match2048 {
				panic("invalid regex match")
			}
		}
		buf.Reset()
	})

	b.Run("match regex 4096", func(b *testing.B) {
		polyglot.Encoder(buf).String(Regex4096)
		b.SetBytes(int64(len(Regex4096)))
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			out, err := plugin.Call("match_regex", buf.Bytes())
			if err != nil {
				panic(err)
			}

			dec := polyglot.GetDecoder(out)
			matches, err := dec.String()
			if err != nil {
				panic(err)
			}
			dec.Return()

			if matches != Match4096 {
				panic("invalid regex match")
			}
		}
		buf.Reset()
	})

	b.Run("match regex 8192", func(b *testing.B) {
		polyglot.Encoder(buf).String(Regex8192)
		b.SetBytes(int64(len(Regex8192)))
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			out, err := plugin.Call("match_regex", buf.Bytes())
			if err != nil {
				panic(err)
			}

			dec := polyglot.GetDecoder(out)
			matches, err := dec.String()
			if err != nil {
				panic(err)
			}
			dec.Return()

			if matches != Match8192 {
				panic("invalid regex match")
			}
		}
		buf.Reset()
	})

	b.Run("match regex 16384", func(b *testing.B) {
		polyglot.Encoder(buf).String(Regex16384)
		b.SetBytes(int64(len(Regex16384)))
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			out, err := plugin.Call("match_regex", buf.Bytes())
			if err != nil {
				panic(err)
			}

			dec := polyglot.GetDecoder(out)
			matches, err := dec.String()
			if err != nil {
				panic(err)
			}
			dec.Return()

			if matches != Match16384 {
				panic("invalid regex match")
			}
		}
		buf.Reset()
	})

	b.Run("match regex 32768", func(b *testing.B) {
		polyglot.Encoder(buf).String(Regex32768)
		b.SetBytes(int64(len(Regex32768)))
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			out, err := plugin.Call("match_regex", buf.Bytes())
			if err != nil {
				panic(err)
			}

			dec := polyglot.GetDecoder(out)
			matches, err := dec.String()
			if err != nil {
				panic(err)
			}
			dec.Return()

			if matches != Match32768 {
				panic("invalid regex match")
			}
		}
		buf.Reset()
	})
}

func BenchmarkNativeGo(b *testing.B) {
	b.Run("match regex 2048", func(b *testing.B) {
		b.SetBytes(int64(len(Regex2048)))
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			matches, err := native.ReplaceAll(Regex2048)
			if err != nil {
				panic(err)
			}

			if matches != Match2048 {
				panic("invalid regex match")
			}
		}
	})

	b.Run("match regex 4096", func(b *testing.B) {
		b.SetBytes(int64(len(Regex4096)))
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			matches, err := native.ReplaceAll(Regex4096)
			if err != nil {
				panic(err)
			}

			if matches != Match4096 {
				panic("invalid regex match")
			}
		}
	})

	b.Run("match regex 8192", func(b *testing.B) {
		b.SetBytes(int64(len(Regex8192)))
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			matches, err := native.ReplaceAll(Regex8192)
			if err != nil {
				panic(err)
			}

			if matches != Match8192 {
				panic("invalid regex match")
			}
		}
	})

	b.Run("match regex 16384", func(b *testing.B) {
		b.SetBytes(int64(len(Regex16384)))
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			matches, err := native.ReplaceAll(Regex16384)
			if err != nil {
				panic(err)
			}

			if matches != Match16384 {
				panic("invalid regex match")
			}
		}
	})

	b.Run("match regex 32768", func(b *testing.B) {
		b.SetBytes(int64(len(Regex32768)))
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			matches, err := native.ReplaceAll(Regex32768)
			if err != nil {
				panic(err)
			}

			if matches != Match32768 {
				panic("invalid regex match")
			}
		}
	})
}
