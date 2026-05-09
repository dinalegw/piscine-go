package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	if len(a) < 2 {
		return true
	}

	isAscending := true
	isDescending := true

	for i := 0; i < len(a)-1; i++ {
		comp := f(a[i], a[i+1])

		if comp > 0 { // a[i] > a[i+1] → violates ascending
			isAscending = false
		}
		if comp < 0 { // a[i] < a[i+1] → violates descending
			isDescending = false
		}
	}

	return isAscending || isDescending
}
