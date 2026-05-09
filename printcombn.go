package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return
	}
	var comb [9]rune
	printCombNRecursive(comb, n, 0, '0')
	z01.PrintRune('\n') // Add newline after all combinations
}

func printCombNRecursive(comb [9]rune, n int, index int, start rune) {
	if index == n {
		for i := 0; i < n; i++ {
			z01.PrintRune(comb[i])
		}
		if comb[0] != rune(10-n)+'0' { // skip comma for last combo
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
		return
	}

	for r := start; r <= '9'; r++ {
		comb[index] = r
		printCombNRecursive(comb, n, index+1, r+1)
	}
}
