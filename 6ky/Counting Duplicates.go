package main

import (
	"fmt"
	"strings"
)

func main() {
	a1 := strings.Split("abcd", "")
	// a2 := strings.Split("abcd", "")
	// count := 0

	for _, x := range a1 {
		fmt.Println(string(x))
	}
}
