package main

import "github.com/01-edu/z01"

func main() {
	for r := '0'; r <= '9'; r++ {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
