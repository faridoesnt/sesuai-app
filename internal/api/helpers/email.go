package helpers

import "regexp"

func IsEmail(text string) bool {
	polaEmail := `\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}\b`

	regexEmail := regexp.MustCompile(polaEmail)

	return regexEmail.MatchString(text)
}
