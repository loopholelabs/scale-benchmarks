package regex

import "regexp"

var (
	r = regexp.MustCompile("p([a-z]+)ch")
)

func FindString(input string) (string, error) {
	return r.FindString(input), nil
}
