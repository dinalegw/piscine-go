package piscine

func Map(f func(int) bool, a []int) []bool {
	result := make([]bool, len(a))
	for i := 0; i < len(a); i++ {
		result[i] = f(a[i])
	}
	return result
}
