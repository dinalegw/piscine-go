package main

import (
	"os"

	"github.com/01-edu/z01"
)

// Bubble sort for string slice
func sortStrings(args []string) {
	for i := 0; i < len(args)-1; i++ {
		for j := 0; j < len(args)-i-1; j++ {
			if args[j] > args[j+1] {
				args[j], args[j+1] = args[j+1], args[j]
			}
		}
	}
}

func main() {
	args := os.Args[1:] // skip program name
	if len(args) == 0 {
		return
	}

	sortStrings(args)

	for _, s := range args {
		for _, r := range s {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}
