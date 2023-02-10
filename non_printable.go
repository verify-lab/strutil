package strutil

import (
	"strings"
	"unicode"
)

func StripNonPrintable(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsPrint(r) {
			return r
		}
		return -1
	}, str)
}
