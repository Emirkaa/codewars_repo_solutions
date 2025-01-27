package main

import "strings"

func is_the_uppercase(vari string) bool {
	formated := strings.ReplaceAll(vari, " ", "")
	formated_2 := strings.ToUpper(formated)
	if formated == formated_2 {
		return true
	} else {
		return false
	}

}
