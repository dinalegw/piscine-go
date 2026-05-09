package main

import (
	"os"

	"github.com/01-edu/z01"
)

func isVowel(r rune) bool {
	switch r {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	case 'A', 'E', 'I', 'O', 'U':
		return true
	}
	return false
}

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		z01.PrintRune('\n')
		return
	}

	// Build full input string with spaces
	input := ""
	for i, arg := range args {
		if i > 0 {
			input += " "
		}
		input += arg
	}

	runes := []rune(input)
	vowelIndices := []int{}

	// collect positions of vowels
	for i, r := range runes {
		if isVowel(r) {
			vowelIndices = append(vowelIndices, i)
		}
	}

	// mirror vowels across the whole string
	for i := 0; i < len(vowelIndices)/2; i++ {
		j := len(vowelIndices) - 1 - i
		runes[vowelIndices[i]], runes[vowelIndices[j]] = runes[vowelIndices[j]], runes[vowelIndices[i]]
	}

	// print result
	for _, r := range runes {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
