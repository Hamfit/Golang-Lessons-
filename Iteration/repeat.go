package iteration

import "strings"

func Repeat(character string, n int) string {
	var repeated strings.Builder
	for range n {
		repeated.WriteString(character)
	}
	return repeated.String()
}