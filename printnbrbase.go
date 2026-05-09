package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	if !isValidBase(base) {
		printString("NV")
		return
	}

	baseLen := uint64(len(base))
	n := int64(nbr)
	var un uint64

	// Handle negative numbers safely
	if n < 0 {
		z01.PrintRune('-')
		// Convert safely using uint64 to avoid overflow on MinInt64
		un = uint64(-(n + 1)) + 1
	} else {
		un = uint64(n)
	}

	// Special case: number = 0
	if un == 0 {
		z01.PrintRune(rune(base[0]))
		return
	}

	// Convert number to base
	var digits []rune
	for un > 0 {
		d := un % baseLen
		digits = append([]rune{rune(base[d])}, digits...)
		un /= baseLen
	}

	// Print digits
	for _, r := range digits {
		z01.PrintRune(r)
	}
}

func isValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}
	seen := make(map[rune]bool)
	for _, r := range base {
		if r == '+' || r == '-' || seen[r] {
			return false
		}
		seen[r] = true
	}
	return true
}

func printString(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}
