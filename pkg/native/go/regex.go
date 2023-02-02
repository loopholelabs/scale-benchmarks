package regex

import "regexp"

func FindString(input string) (string, error) {
	r, err := regexp.Compile("p([a-z]+)ch")
	if err != nil {
		return "", err
	}

	return r.FindString(input), nil
}
