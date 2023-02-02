//go:build tinygo || js || wasm

package scale

import (
	signature "github.com/loopholelabs/scale-benchmarks/pkg/go/signature/bad-signature"
)

func Scale(ctx *signature.Context) (*signature.Context, error) {
	ctx.Data = 30

	return ctx, nil
}
