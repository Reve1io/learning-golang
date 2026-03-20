package generatingstringsinloops

import (
	"strings"
	"unicode"
)

func FilterString(str string, i rune) string {
	var bld strings.Builder

	if len(str) == 0 {
		return ""
	}

	if i == 0 {
		return str
	}

	i = unicode.ToLower(i)

	for _, R := range str {

		if R == '\u0020' {
			bld.WriteRune(R)
		} else if unicode.IsUpper(R) {
			r := unicode.ToLower(R)

			if r != i {
				lowR := unicode.ToUpper(r)
				bld.WriteRune(lowR)
			}
		} else if unicode.IsLower(R) {

			if R != i {
				bld.WriteRune(R)
			}
		}
	}

	result := bld.String()
	return result
}
