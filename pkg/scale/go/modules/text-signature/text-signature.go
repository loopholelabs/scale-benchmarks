//go:build tinygo || js || wasm

package scale

import (
	"regexp"

	signature "github.com/loopholelabs/scale-benchmarks/pkg/scale/go/signature/text-signature"
)

var (
	r = regexp.MustCompile(`\b\w{4}\b`)
)

func Scale(ctx *signature.Context) (*signature.Context, error) {
	ctx.Data = r.ReplaceAllString(ctx.Data, "wasm")
	return ctx, nil
}
