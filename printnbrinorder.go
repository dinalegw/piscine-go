package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n < 0 {
		return
	}

	digits := [10]int{}
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	// Count digits
	for n > 0 {
		d := n % 10
		digits[d]++
		n /= 10
	}

	// Print digits in ascending order
	for i := 0; i <= 9; i++ {
		for j := 0; j < digits[i]; j++ {
			z01.PrintRune(rune('0' + i))
		}
	}
}
