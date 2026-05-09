package piscine

func Atoi(s string) int {
	if len(s) == 0 {
		return 0
	}

	sign := 1
	start := 0

	// Handle sign
	if s[0] == '-' {
		sign = -1
		start = 1
	} else if s[0] == '+' {
		start = 1
	}

	// If there is another + or - OR the string is just "+" or "-"
	if start == 1 && len(s) == 1 {
		return 0
	}
	if start == 1 && (s[1] == '+' || s[1] == '-') {
		return 0
	}

	result := 0

	for i := start; i < len(s); i++ {
		ch := s[i]
		if ch < '0' || ch > '9' {
			return 0
		}
		result = result*10 + int(ch-'0')
	}

	return result * sign
}
