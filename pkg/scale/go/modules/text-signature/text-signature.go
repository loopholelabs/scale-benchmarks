//go:build tinygo || js || wasm

package scale

import (
	"regexp"

	signature "github.com/loopholelabs/scale-benchmarks/pkg/scale/go/signature/text-signature"
)

func Scale(ctx *signature.Context) (*signature.Context, error) {
	r, err := regexp.Compile("p([a-z]+)ch")
	if err != nil {
		return ctx, err
	}

	ctx.Data = r.FindString(ctx.Data)

	return ctx, nil
}
