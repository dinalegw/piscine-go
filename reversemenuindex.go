package piscine

func ReverseMenuIndex(menu []string) []string {
	n := len(menu)
	reversed := make([]string, n)

	for i := 0; i < n; i++ {
		reversed[i] = menu[n-1-i]
	}

	return reversed
}
