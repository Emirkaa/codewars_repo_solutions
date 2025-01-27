package kata

import "strings"

func Contamination(text, char string) string {
	for _, charrr := range text {
		text = strings.ReplaceAll(text, string(charrr), string(char))
	}
	return text
}
package kata


