package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(spin("Hey fellow warriors"))
}

func polindrome(sh string) string {
	res := ""
	for _, x := range sh {
		res = string(x) + res

	}
	return res
}

func spin(s string) string {

	spl := strings.Split(s, " ")
	result := ""
	for _, x := range spl {
		if len(x) < 5 {
			result += x + " "
		} else {
			result += polindrome(x) + " "
		}

	}
	return strings.TrimSpace(result)
}
