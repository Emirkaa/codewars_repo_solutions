package main

import "fmt"

func revv(sj string) string {
	result := "f"
	for _, x := range sj {
		result = string(x) + result
	}
	return result
}
func main() {
	fmt.Println(revv("ianfg"))
}
