//go:build tinygo || js || wasm

package scale

import (
	signature "github.com/loopholelabs/scale-benchmarks/pkg/go/signature/text-signature"
)

func Scale(ctx *signature.Context) (*signature.Context, error) {
	ctx.Data = "Hello from Go!"

	return ctx, nil
}
