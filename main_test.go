package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"testing"

	"github.com/extism/extism"
	regex "github.com/loopholelabs/scale-benchmarks/pkg/native/go"
	"github.com/loopholelabs/scale-benchmarks/pkg/scale/go/signature/text-signature"
	runtime "github.com/loopholelabs/scale/go"
	"github.com/loopholelabs/scale/go/tests/harness"
	"github.com/loopholelabs/scalefile"
	"github.com/loopholelabs/scalefile/scalefunc"
)

type existmOutput struct {
	Matches string `json:"matches"`
}

func Test_scale_go(t *testing.T) {
	moduleConfig := &harness.Module{
		Name:      "text-signature",
		Path:      "pkg/scale/go/modules/text-signature/text-signature.go",
		Signature: "github.com/loopholelabs/scale-benchmarks/pkg/scale/go/signature/text-signature",
	}

	generatedModules := harness.GoSetup(
		t,
		[]*harness.Module{moduleConfig},
		"github.com/loopholelabs/scale-benchmarks/pkg/scale/go/modules",
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
		Language:  scalefunc.Go,
		Function:  module,
	}

	r, err := runtime.NewWithSignature(context.Background(), text.New, []*scalefunc.ScaleFunc{scaleFunc})
	if err != nil {
		panic(err)
	}

	i, err := r.Instance(nil)
	if err != nil {
		panic(err)
	}

	i.Context().Data = "peach"

	if err := i.Run(context.Background()); err != nil {
		panic(err)
	}

	log.Println("Scale Go:", i.Context().Data)
}

func Test_scale_rust(t *testing.T) {
	moduleConfig := &harness.Module{
		Name:          "text_signature",
		Path:          "./pkg/scale/rust/modules/text_signature/text_signature.rs",
		Signature:     "text_signature",
		SignaturePath: "../../../signature/text-signature",
	}

	generatedModules := harness.RustSetup(
		t,
		[]*harness.Module{moduleConfig},
		[]*scalefile.Dependency{
			{
				Name:    "scale_signature",
				Version: "0.2.0",
			},
			{
				Name:    "wee_alloc",
				Version: "0.4.5",
			},
			{
				Name:    "regex",
				Version: "1.7.1",
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

	r, err := runtime.NewWithSignature(context.Background(), text.New, []*scalefunc.ScaleFunc{scaleFunc})
	if err != nil {
		panic(err)
	}

	i, err := r.Instance(nil)
	if err != nil {
		panic(err)
	}

	i.Context().Data = "peach"

	if err := i.Run(context.Background()); err != nil {
		panic(err)
	}

	log.Println("Scale Rust:", i.Context().Data)
}

func Test_native_go(t *testing.T) {
	matches, err := regex.FindString("peach")
	if err != nil {
		panic(err)
	}

	log.Println("Native Go:", matches)
}

func Test_exism_rust(t *testing.T) {
	cmd := exec.Command("cargo", "build", "--release", "--target", "wasm32-unknown-unknown")

	cmd.Dir = "./pkg/extism/rust"

	if err := cmd.Run(); err != nil {
		panic(err)
	}

	ctx := extism.NewContext()
	defer ctx.Free()

	manifest := extism.Manifest{Wasm: []extism.Wasm{extism.WasmFile{Path: "./pkg/extism/rust/target/wasm32-unknown-unknown/release/rust.wasm"}}}

	plugin, err := ctx.PluginFromManifest(manifest, []extism.Function{}, true)
	if err != nil {
		panic(err)
	}

	out, err := plugin.Call("match_regex", []byte("peach"))
	if err != nil {
		panic(err)
	}

	var dst existmOutput
	if json.Unmarshal(out, &dst); err != nil {
		panic(err)
	}

	fmt.Println("Exism Rust:", dst.Matches)
}
