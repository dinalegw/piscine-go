package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// os.Args[0] is the program name, so start from 1
	for i := 1; i < len(os.Args); i++ {
		for _, r := range os.Args[i] {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}
