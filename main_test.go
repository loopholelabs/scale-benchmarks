package main

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/loopholelabs/scale-benchmarks/pkg/go/signature/bad-signature"
	runtime "github.com/loopholelabs/scale/go"
	"github.com/loopholelabs/scale/go/tests/harness"
	"github.com/loopholelabs/scalefile"
	"github.com/loopholelabs/scalefile/scalefunc"
)

func Test_main(t *testing.T) {
	{
		moduleConfig := &harness.Module{
			Name:      "bad-signature",
			Path:      "pkg/go/modules/bad-signature/bad-signature.go",
			Signature: "github.com/loopholelabs/scale-benchmarks/pkg/go/signature/bad-signature",
		}

		generatedModules := harness.GoSetup(
			t,
			[]*harness.Module{moduleConfig},
			"github.com/loopholelabs/scale-benchmarks/pkg/go/modules",
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

		r, err := runtime.NewWithSignature(context.Background(), bad.New, []*scalefunc.ScaleFunc{scaleFunc})
		if err != nil {
			panic(err)
		}

		i, err := r.Instance(nil)
		if err != nil {
			panic(err)
		}

		if err := i.Run(context.Background()); err != nil {
			panic(err)
		}

		log.Println("Go:", i.Context().Data)
	}

	{
		moduleConfig := &harness.Module{
			Name:          "bad_signature",
			Path:          "./pkg/rust/modules/bad_signature/bad_signature.rs",
			Signature:     "bad_signature",
			SignaturePath: "../../../signature/bad-signature",
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

		r, err := runtime.NewWithSignature(context.Background(), bad.New, []*scalefunc.ScaleFunc{scaleFunc})
		if err != nil {
			panic(err)
		}

		i, err := r.Instance(nil)
		if err != nil {
			panic(err)
		}

		if err := i.Run(context.Background()); err != nil {
			panic(err)
		}

		log.Println("Rust:", i.Context().Data)
	}
}
