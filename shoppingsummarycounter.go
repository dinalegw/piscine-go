package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	result := make(map[string]int)
	word := ""

	for i := 0; i < len(str); i++ {
		c := str[i]
		if c == ' ' { // split only on the single space character
			// append current token (may be empty) and reset
			result[word]++
			word = ""
		} else {
			word += string(c)
		}
	}
	// add the final token (may be empty if string ends with space)
	result[word]++

	return result
}
