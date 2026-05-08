package main

import (
	"fmt"
	"strings"
)

func fixquote(text string) string {
	text = strings.ReplaceAll(text, "' ", "'")
	text = strings.ReplaceAll(text, " '", "'")
	return text
}

func main() {
	fmt.Println(fixquote("' awesome '"))
	fmt.Println(fixquote("' hello world '"))
}