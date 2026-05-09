package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printHelp() {
	helpText := "--insert\n  -i\n\t This flag inserts the string into the string passed as argument.\n--order\n  -o\n\t This flag will behave like a boolean, if it is called it will order the argument."
	for _, r := range helpText {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func sortRunes(r []rune) {
	n := len(r)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if r[j] > r[j+1] {
				r[j], r[j+1] = r[j+1], r[j]
			}
		}
	}
}

func hasPrefix(s, prefix string) bool {
	if len(s) < len(prefix) {
		return false
	}
	for i := 0; i < len(prefix); i++ {
		if s[i] != prefix[i] {
			return false
		}
	}
	return true
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		printHelp()
		return
	}

	insertStr := ""
	order := false
	argStr := ""
	skipNext := false

	for i := 0; i < len(args); i++ {
		if skipNext {
			skipNext = false
			continue
		}
		arg := args[i]

		if arg == "--help" || arg == "-h" {
			printHelp()
			return
		} else if hasPrefix(arg, "--insert=") {
			insertStr = arg[len("--insert="):]
		} else if hasPrefix(arg, "-i=") {
			insertStr = arg[len("-i="):]
		} else if arg == "-i" {
			if i+1 < len(args) {
				insertStr = args[i+1]
				skipNext = true
			}
		} else if arg == "--order" || arg == "-o" {
			order = true
		} else if len(arg) > 0 && arg[0] != '-' {
			argStr = arg
		}
	}

	result := argStr + insertStr

	if order {
		runes := []rune(result)
		sortRunes(runes)
		result = string(runes)
	}

	for _, r := range result {
		z01.PrintRune(r)
	}

	if len(result) > 0 {
		z01.PrintRune('\n')
	}
}
