package piscine

func Compact(ptr *[]string) int {
	if ptr == nil {
		return 0
	}

	s := *ptr
	n := len(s)
	// allocate exactly once using make (allowed)
	out := make([]string, n)

	j := 0
	for i := 0; i < n; i++ {
		if s[i] != "" {
			out[j] = s[i]
			j++
		}
	}

	*ptr = out[:j]
	return j
}
