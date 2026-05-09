package piscine

import "github.com/01-edu/z01"

func DescendComb() {
	first := true

	for a := 99; a >= 0; a-- {
		for b := a - 1; b >= 0; b-- {
			if !first {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
			first = false

			z01.PrintRune(rune(a/10 + '0'))
			z01.PrintRune(rune(a%10 + '0'))
			z01.PrintRune(' ')
			z01.PrintRune(rune(b/10 + '0'))
			z01.PrintRune(rune(b%10 + '0'))
		}
	}
}
