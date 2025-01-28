package kata

func Solve(arr []int) int {
	for _, n := range arr {
		flag := false
		m := -n
		for _, number := range arr {
			if number == m {
				flag = true
			}
		}
		if !flag {
			return n
		}
	}
	return 0
}
