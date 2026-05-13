package main 

import (
	"fmt"
	"strings"
)

func ContainsAny(s, chars string) bool {
	return strings.ContainsAny(s, chars)
}

func main() {
	fmt.Println(ContainsAny("hello", "@#$"))
	fmt.Println(ContainsAny("hello@", "@#$"))
	fmt.Println(ContainsAny("team", "i"))
	fmt.Println(ContainsAny("fail", "ui"))
	fmt.Println(ContainsAny("ure", "ui"))
	fmt.Println(ContainsAny("foo", ""))
	fmt.Println(ContainsAny("", ""))
}