package kata

import "strconv"

func DigitalRoot(n int) int {

	for n >= 10 {
		strN := strconv.Itoa(n)
		g := 0
		for _, x := range strN {

			num, _ := strconv.Atoi(string(x))
			g += num

		}
		n = g

	}
	return n

}
