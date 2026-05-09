package main

import (
	"os"

	"github.com/01-edu/z01"
)

// manual atoi: convert string to int
func atoi(s string) (int, bool) {
	if len(s) == 0 {
		return 0, false
	}
	n := 0
	for _, c := range s {
		if c < '0' || c > '9' {
			return 0, false
		}
		n = n*10 + int(c-'0')
	}
	return n, true
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		// no arguments → nothing to print
		return
	}

	upper := false

	// check if first argument is --upper
	if args[0] == "--upper" {
		upper = true
		args = args[1:]
		if len(args) == 0 {
			// flag only, no letters to print → nothing
			return
		}
	}

	for _, arg := range args {
		n, ok := atoi(arg)
		if !ok || n < 1 || n > 26 {
			z01.PrintRune(' ')
			continue
		}
		if upper {
			z01.PrintRune(rune('A' + n - 1))
		} else {
			z01.PrintRune(rune('a' + n - 1))
		}
	}

	// print newline only if there was at least one argument
	z01.PrintRune('\n')
}
