package main

import "strings"

type MyString string

func main() {
	strrr := "o gosfj  OJFoj foe ofijw"
	is_the_uppercase(strrr)

}

func (s MyString) is_the_uppercase() bool {
	formated := strings.ReplaceAll(string(s), " ", "")
	return formated == strings.ToUpper(formated)
}
