package main

import (
	"bufio"
	"os"

	"github.com/01-edu/z01"
)

func printString(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		printString("ERROR: open " + filename + ": no such file or directory\n")
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		printString(scanner.Text())
		z01.PrintRune('\n')
	}
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		// Read from stdin
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			printString(scanner.Text())
			z01.PrintRune('\n')
		}
		return
	}

	for _, filename := range args {
		printFile(filename) // exits immediately if file cannot be opened
	}
}
