package main

import (
	"fmt"
	"strings"
	"unicode"
)

func ToJadenCase(str string) string {
	curent_word := ""

	var words []string
	for _, y := range str {
		if (y >= 'a' && y <= 'z') || (y >= 'A' && y <= 'Z') {
			curent_word += string(y)
		} else if curent_word != "" {
			words = append(words, curent_word)
			curent_word = ""
		}

	}
	var result []string

	for _, word := range words {
		// Если слово не пустое, преобразуем первую букву в заглавную
		if len(word) > 0 {
			// Преобразуем первую букву в заглавную, а остальные — в строчные
			runes := []rune(word)
			runes[0] = unicode.ToUpper(runes[0])
			for i := 1; i < len(runes); i++ {
				runes[i] = unicode.ToLower(runes[i])
			}
			result = append(result, string(runes))
		}
	}
	return strings.Join(result, " ")

}

func main() {
	sadfj := "asfjh asd JHFDKSFH  kjSfhsjhkdf"
	fmt.Println(ToJadenCase(sadfj))
}

/*
package kata

import (
	"strings"
	"unicode"
)

func ToJadenCase(str string) string {
	curent_word := ""
	var words []string

	// Перебираем все символы строки
	for _, y := range str {
		if (y >= 'a' && y <= 'z') || (y >= 'A' && y <= 'Z') {
			// Если символ — буква, добавляем в текущее слово
			curent_word += string(y)
		} else if curent_word != "" {
			// Если слово завершено (встретили пробел), добавляем его в массив
			words = append(words, curent_word)
			curent_word = ""
		}
	}

	// Обработка последнего слова (если оно осталось)
	if curent_word != "" {
		words = append(words, curent_word)
	}

	// Преобразуем каждое слово в JadenCase
	var result []string
	for _, word := range words {
		if len(word) > 0 {
			// Преобразуем первую букву в заглавную
			runes := []rune(word)
			runes[0] = unicode.ToUpper(runes[0])
			// Преобразуем остальные буквы в строчные
			for i := 1; i < len(runes); i++ {
				runes[i] = unicode.ToLower(runes[i])
			}
			result = append(result, string(runes))
		}
	}

	// Собираем все слова в одну строку
	return strings.Join(result, " ")
}
\*
