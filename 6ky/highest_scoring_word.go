package kata

import "strings"

func f_b_w_s(s string) int {
	score := 0

	for _, x := range s {
		score += int(x - 'a' + 1)
	}
	return score
}

func High(s string) string {
	max_word := ""
	max_score := 0
	reform_srt := strings.Split(s, " ")
	for _, x := range reform_srt {
		f := f_b_w_s(x)
		if f > max_score {
			max_word = string(x)
			max_score = f
		}
	}
	return max_word
}
