package piscine

func LastRune(s string) rune {
	var last rune
	for _, r := range s {
		last = r
	}
	return last
}
