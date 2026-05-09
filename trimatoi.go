package piscine

func TrimAtoi(s string) int {
	sign := 1
	num := 0
	signApplied := false

	for _, r := range s {
		if r == '-' && !signApplied && num == 0 {
			sign = -1
			signApplied = true
		} else if r >= '0' && r <= '9' {
			num = num*10 + int(r-'0')
		}
	}
	return num * sign
}
