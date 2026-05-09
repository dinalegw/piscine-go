package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	indexInBase := func(c rune, base string) int {
		for i, b := range base {
			if b == c {
				return i
			}
		}
		return -1
	}
	value := 0
	baseFromLen := len(baseFrom)
	for _, c := range nbr {
		value = value*baseFromLen + indexInBase(c, baseFrom)
	}

	if value == 0 {
		return string(baseTo[0])
	}

	baseToLen := len(baseTo)
	result := ""
	for value > 0 {
		result = string(baseTo[value%baseToLen]) + result
		value /= baseToLen
	}
	return result
}
