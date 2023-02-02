//go:build tinygo || js || wasm

package scale

import (
	"regexp"

	signature "github.com/loopholelabs/scale-benchmarks/pkg/scale/go/signature/text-signature"
)

var (
	r = regexp.MustCompile("p([a-z]+)ch")
)

func Scale(ctx *signature.Context) (*signature.Context, error) {
	ctx.Data = r.FindString(ctx.Data)
	return ctx, nil
}
