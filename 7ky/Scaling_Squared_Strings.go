package kata

import (
	"fmt"
	"strings"
)

func Scale(s string, k, n int) (result string) {
	subStr := strings.Split(s, "\n")
	for _, value := range subStr {
		var hscale string
		for _, char := range value {
			hscale += strings.Repeat(string(char), k)
		}

		result += strings.Repeat(hscale+"\n", n)
	}
	return strings.TrimRight(result, "\n")
}
func main() {
	fmt.Println(Scale("\npoausfp\nianrgp nawrg |\nopahrgp norgg\n", 2, 2))

}
