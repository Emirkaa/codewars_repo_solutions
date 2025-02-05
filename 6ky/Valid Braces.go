package main

func ValidBraces(str string) bool {
	mnoj1 := map[string]struct{}{
		"{": {},
		"(": {},
		"[": {},
	}

	mnoj2 := map[string]struct{}{
		"}": {},
		"]": {},
		")": {},
	}

	half1 := str[:len(str)/2]
	half2 := str[len(str)/2:]

	count1 := 0
	count2 := 0

	lenn1 := len(half1)
	lenn2 := len(half2)

	for _, x := range half1 {
		if _, exist := mnoj1[string(x)]; exist {
			count1 += 1
		}
	}

	for _, y := range half2 {
		if _, exist2 := mnoj2[string(y)]; exist2 {
			count2 += 1
		}
	}
	if count1 == lenn1 && count2 == lenn2 {
		return true
	}
	return false
}
