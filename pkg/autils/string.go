package autils

import "strings"

func UncapitalizedInitialLetter(s string) string {
	firstChar := strings.ToLower(s[0:1])

	return firstChar + s[1:]
}
