package kata

import "strings"

func contamination(texttt, charrr string) string {
	return strings.Repeat(charrr, len(texttt))
}
