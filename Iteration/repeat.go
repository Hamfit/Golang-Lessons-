package iteration

import "strings"

const repeatCount = 5

func Repeat(character string) string {
	var repeated strings.Builder
	for range 5 {
		repeated .WriteString(character)
	}
	return repeated.String()
}