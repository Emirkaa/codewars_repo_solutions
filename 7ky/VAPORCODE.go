package kata

import "strings"

func Vaporcode(s string) string {
	ss := ""
	for _, x := range s {
		if string(x) != " " {
			ss += strings.ToUpper(string(x)) + "  "
		}

	}

	return strings.TrimSpace(ss)
}
