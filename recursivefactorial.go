package piscine

func RecursiveFactorial(nb int) int {
	if nb < 0 || nb > 20 { // invalid numbers or too large
		return 0
	}
	if nb == 0 || nb == 1 {
		return 1
	}
	return nb * RecursiveFactorial(nb-1)
}
