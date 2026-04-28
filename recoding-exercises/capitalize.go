// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func Capitalize(s string) string {
// 	words := strings.Fields(s)
// 	for i, char := range words {
// 		if len(char) > 0 {
// 			words[i] = strings.ToUpper(string(char[0])) + strings.ToLower(char[1:])
// 		}
// 	}
// 	return strings.Join(words, " ")
// }

// func main() {
// 	str := "hello world! My name is james bond"
// 	fmt.Println(Capitalize(str))
// }
