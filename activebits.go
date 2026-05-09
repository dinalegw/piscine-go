package piscine

func ActiveBits(n int) int {
	count := 0
	for n != 0 {
		// Brian Kernighan’s algorithm: removes the lowest active bit
		n = n & (n - 1)
		count++
	}
	return count
}
