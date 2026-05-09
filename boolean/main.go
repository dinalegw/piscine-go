package main

import (
	"os"

	"github.com/01-edu/z01"
)

const (
	EvenMsg = "I have an even number of arguments"
	OddMsg  = "I have an odd number of arguments"

	yes = true
	no  = false
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func even(nbr int) bool {
	return nbr%2 == 0
}

func isEven(nbr int) bool {
	if even(nbr) {
		return yes
	}
	return no
}

func main() {
	lengthOfArg := len(os.Args) - 1

	if isEven(lengthOfArg) {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}
