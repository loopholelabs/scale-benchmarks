package regex

import "regexp"

var (
	r = regexp.MustCompile(`\b\w{4}\b`)
)

func FindString(input string) (string, error) {
	return r.ReplaceAllString(input, "wasm"), nil
}
