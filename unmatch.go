package piscine

func Unmatch(a []int) int {
	m := make(map[int]int)

	for _, n := range a {
		m[n]++
	}

	for _, n := range a {
		if m[n]%2 == 1 {
			return n
		}
	}
	return -1
}
