package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// start from the last argument and go backwards
	for i := len(os.Args) - 1; i > 0; i-- {
		for _, r := range os.Args[i] {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}
