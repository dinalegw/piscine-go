package piscine

func NRune(s string, n int) rune {
	if n <= 0 {
		return 0
	}
	i := 1
	for _, r := range s {
		if i == n {
			return r
		}
		i++
	}
	return 0
}
