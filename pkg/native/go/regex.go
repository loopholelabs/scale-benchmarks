package regex

import "regexp"

var (
	r = regexp.MustCompile("peach")
)

func FindString(input string) (string, error) {
	return r.FindString(input), nil
}
