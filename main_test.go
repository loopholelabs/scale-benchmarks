package main

import (
	"context"
	"os"
	"os/exec"
	"testing"

	"github.com/extism/extism"
	"github.com/loopholelabs/polyglot-go"
	regex "github.com/loopholelabs/scale-benchmarks/pkg/native/go"
	"github.com/loopholelabs/scale-benchmarks/pkg/scale/go/signature/text-signature"
	runtime "github.com/loopholelabs/scale/go"
	"github.com/loopholelabs/scale/go/tests/harness"
	"github.com/loopholelabs/scalefile"
	"github.com/loopholelabs/scalefile/scalefunc"
)

const (
	Regex = "BYwnOjWhprPmDncp8qpQ5CY4r1RGZuqKLBowmtMCd test ETjLOG685YC4RIjXB0HadNpqYS4M7GPGUVAKRZRC1ibqQqGnuzqX2Hjosm6MKNCp5QifX7Up2phqkFqkjpSu3k59oi6M5YbTMiy4JukVFx2402IlrHU1McK7US0skB1cF0W2ZDpsypNmGJRXRMY0pPsYbw7G2a0xJnhTITXcuF5xJWR1rz5zdGZQbbjZoHZcEnveDFq5kOmCVc test DsJVHTsAlypLI9sVtbTLwmE1DG2C6AgUo3GO1DpCx3jV43oXUxaTVJqZO13AYqvNPbxizYZ5BckZFBbJybY3Vnm20Sm7nXbwZs5N2ugz3EpUQvXwqHdHWzc1T8uKPD5LTDM8UBpVoF test 9G3mWarrp43SvoidITriFhzHmyVWNd6n2LIVocr3pOai4DOlkAn7QDup6z6spMAf8UcI4wbfoSzG0k5Qy1rGBhPaJKJRW2 MC9ma3U3rnjAOBtEUHZ2qfOUpfMNgPlGpvzr4IGNNFf9RFlF7yRUBvRnYxyonIWPPiR1x1wWgxc20o5cW4GU7kytAOuGlpzpykcAxCJLLP6wJegaMhAeb8xBLpuBetNEbfcyyOcJBun5BhmFOmv8 test IvICWx2wlYZ61YDBpPcIpqnMb9MHwT8GroC1YITZBlNGBHMpAe4d2sNZe9d0Wvfbv5mMo30Bm1Pa5S3x38jgu6y0BaqZl9GhlukE9CqPJGUsJZ5suDH19WiOrvz7mXwXhi4lWm1YdwNi0xhVnXITtmKq5rikIS6dul1USgDf3TwyLYpyCG46Xj92PssJmnhPdH1WAnvXY sbs8RaemyqmPggtGNwU2JjuPjdmQRakIusv2WimN7zG8R8Pf1225IAJ2j8aiZBrxnjmrucaYOQCrLm7e2Q5q8 test HOkCEJJGHVLYJtGgHKa1PRQ5qCcsIAUdkW3yRfdulutteLe3We9z9XQvWuTYMLDPpOJqMzDNTGpTYts7AL8pFog1k82XVuMZ6ItccxOBpuzDcahH4wDqCGjak8qPVxmnrGmSsrdUHVz6SrScElMo0nOF8RIpYAVdJr5NxWIK1uzc1iIiZnbUD6uDNmBkmfec6IgK6aqnEZaGLDJXDHSYfzWUOi7y3KNPl0CghL9BId8v4040mCKMfmdthWWLJ2tpWIo1482ghiU5 2qtrzgFgYKfyfr4X6FXzN3hM3bLnuwItQrTCEp3BYz79bCAaQGhicZzqE83Mh2 test IIVID622qlEyVEGuEmNJ5JteEzbpklhTKnVMflzzWyWbZe6kIgeUr9mxWjkJGisvRbZKwfnojeC82M1nHgUa4k46x7Dw7mL3rChORjBxBMYjFeOvCsT6kEo3vPeachLUKdkExJbr9Yei0fKyOFSDlxpFhlRKuwGxXu4jGo4CzKDsVsahqzC9iGw53bHiw0V4Pwmdhzv482s3zU9XLTgQr6GuL1I0kSfh9BkVoK5fFvg1hm7ECrt6p8q3kLVjxte EK9W9q2q9etMaPLymcCRZ0XauMDzJY08JeVvovnT2g5hxE7UGW1 test YRotQUivrrXQnhEw55faznZZBU1ULVs4BfYkIkEfS91NetBhona6zrzDwMsXi0FJjdaiJ25lvetPDaMzUs0l6nfkGkVyU376mFPfPkpBKZR2z2Xwzxndi0SkUnqm8jCa7iq2oSJstTdUXtCK2xTXMIh7tiuPVftit GFYQXXI3vY QFe1xShWJgFAqYguQ8gcxMPSzMlyDaPmMuTPgFZDM0cd test NS3fTggxBa4p5jgS4S0nhae05RkYkXGzuNMXeu6IoR9PFqVFnXcBYD0Ld9otrAiqUuIGYGmAjm3WxpZjXfbF2HhWcx86BaSPTut4OwkspEdRCEm1iAzEGJyEOFi93u9eTGpqrKvU09ZcNtO9138cnva2UtIFaRhL02ckRfycz3BGfwqYl3TGtjWdKjmxn1WreRIIq5gkbWJws5VQsov0V2U8pGedj N2RDqWgh2tFiJA9fmytgRgqSnqxIwyBMgY5RnE6CZ0 test Iv4QPiWMu0oG7r0e4nSNtG13O test"
	Match = "BYwnOjWhprPmDncp8qpQ5CY4r1RGZuqKLBowmtMCd wasm ETjLOG685YC4RIjXB0HadNpqYS4M7GPGUVAKRZRC1ibqQqGnuzqX2Hjosm6MKNCp5QifX7Up2phqkFqkjpSu3k59oi6M5YbTMiy4JukVFx2402IlrHU1McK7US0skB1cF0W2ZDpsypNmGJRXRMY0pPsYbw7G2a0xJnhTITXcuF5xJWR1rz5zdGZQbbjZoHZcEnveDFq5kOmCVc wasm DsJVHTsAlypLI9sVtbTLwmE1DG2C6AgUo3GO1DpCx3jV43oXUxaTVJqZO13AYqvNPbxizYZ5BckZFBbJybY3Vnm20Sm7nXbwZs5N2ugz3EpUQvXwqHdHWzc1T8uKPD5LTDM8UBpVoF wasm 9G3mWarrp43SvoidITriFhzHmyVWNd6n2LIVocr3pOai4DOlkAn7QDup6z6spMAf8UcI4wbfoSzG0k5Qy1rGBhPaJKJRW2 MC9ma3U3rnjAOBtEUHZ2qfOUpfMNgPlGpvzr4IGNNFf9RFlF7yRUBvRnYxyonIWPPiR1x1wWgxc20o5cW4GU7kytAOuGlpzpykcAxCJLLP6wJegaMhAeb8xBLpuBetNEbfcyyOcJBun5BhmFOmv8 wasm IvICWx2wlYZ61YDBpPcIpqnMb9MHwT8GroC1YITZBlNGBHMpAe4d2sNZe9d0Wvfbv5mMo30Bm1Pa5S3x38jgu6y0BaqZl9GhlukE9CqPJGUsJZ5suDH19WiOrvz7mXwXhi4lWm1YdwNi0xhVnXITtmKq5rikIS6dul1USgDf3TwyLYpyCG46Xj92PssJmnhPdH1WAnvXY sbs8RaemyqmPggtGNwU2JjuPjdmQRakIusv2WimN7zG8R8Pf1225IAJ2j8aiZBrxnjmrucaYOQCrLm7e2Q5q8 wasm HOkCEJJGHVLYJtGgHKa1PRQ5qCcsIAUdkW3yRfdulutteLe3We9z9XQvWuTYMLDPpOJqMzDNTGpTYts7AL8pFog1k82XVuMZ6ItccxOBpuzDcahH4wDqCGjak8qPVxmnrGmSsrdUHVz6SrScElMo0nOF8RIpYAVdJr5NxWIK1uzc1iIiZnbUD6uDNmBkmfec6IgK6aqnEZaGLDJXDHSYfzWUOi7y3KNPl0CghL9BId8v4040mCKMfmdthWWLJ2tpWIo1482ghiU5 2qtrzgFgYKfyfr4X6FXzN3hM3bLnuwItQrTCEp3BYz79bCAaQGhicZzqE83Mh2 wasm IIVID622qlEyVEGuEmNJ5JteEzbpklhTKnVMflzzWyWbZe6kIgeUr9mxWjkJGisvRbZKwfnojeC82M1nHgUa4k46x7Dw7mL3rChORjBxBMYjFeOvCsT6kEo3vPeachLUKdkExJbr9Yei0fKyOFSDlxpFhlRKuwGxXu4jGo4CzKDsVsahqzC9iGw53bHiw0V4Pwmdhzv482s3zU9XLTgQr6GuL1I0kSfh9BkVoK5fFvg1hm7ECrt6p8q3kLVjxte EK9W9q2q9etMaPLymcCRZ0XauMDzJY08JeVvovnT2g5hxE7UGW1 wasm YRotQUivrrXQnhEw55faznZZBU1ULVs4BfYkIkEfS91NetBhona6zrzDwMsXi0FJjdaiJ25lvetPDaMzUs0l6nfkGkVyU376mFPfPkpBKZR2z2Xwzxndi0SkUnqm8jCa7iq2oSJstTdUXtCK2xTXMIh7tiuPVftit GFYQXXI3vY QFe1xShWJgFAqYguQ8gcxMPSzMlyDaPmMuTPgFZDM0cd wasm NS3fTggxBa4p5jgS4S0nhae05RkYkXGzuNMXeu6IoR9PFqVFnXcBYD0Ld9otrAiqUuIGYGmAjm3WxpZjXfbF2HhWcx86BaSPTut4OwkspEdRCEm1iAzEGJyEOFi93u9eTGpqrKvU09ZcNtO9138cnva2UtIFaRhL02ckRfycz3BGfwqYl3TGtjWdKjmxn1WreRIIq5gkbWJws5VQsov0V2U8pGedj N2RDqWgh2tFiJA9fmytgRgqSnqxIwyBMgY5RnE6CZ0 wasm Iv4QPiWMu0oG7r0e4nSNtG13O wasm"
)

func BenchmarkScaleRust(b *testing.B) {
	moduleConfig := &harness.Module{
		Name:          "text_signature",
		Path:          "./pkg/scale/rust/modules/text_signature/text_signature.rs",
		Signature:     "text_signature",
		SignaturePath: "../../../signature/text-signature",
	}

	generatedModules := RustSetup(
		b,
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

	r, err := runtime.NewWithSignature(context.Background(), text.New, []*scalefunc.ScaleFunc{scaleFunc})
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	b.ResetTimer()
	b.Run("match_regex", func(b *testing.B) {
		b.SetBytes(int64(len(Regex)))
		b.ReportAllocs()
		for j := 0; j < b.N; j++ {
			i, err := r.Instance(nil)
			if err != nil {
				panic(err)
			}

			i.Context().Data = Regex

			if err := i.Run(ctx); err != nil {
				panic(err)
			}

			if i.Context().Data != Match {
				panic("invalid regex match")
			}
		}
	})
}

func BenchmarkExtismRust(b *testing.B) {
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

	buf := polyglot.NewBuffer()
	polyglot.Encoder(buf).String(Regex)

	b.ResetTimer()
	b.Run("match_regex", func(b *testing.B) {
		b.SetBytes(int64(len(Regex)))
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

			if matches != Match {
				panic("invalid regex match")
			}
		}
	})
}

func BenchmarkNativeGo(b *testing.B) {
	b.Run("match_regex", func(b *testing.B) {
		b.SetBytes(int64(len(Regex)))
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			matches, err := regex.FindString(Regex)
			if err != nil {
				panic(err)
			}

			if matches != Match {
				panic("invalid regex match")
			}
		}
	})
}
