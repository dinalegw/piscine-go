package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "usage: ztail -c N file...")
		os.Exit(1)
	}

	if os.Args[1] != "-c" {
		fmt.Fprintln(os.Stderr, "only -c option is supported")
		os.Exit(1)
	}

	var n int
	_, err := fmt.Sscanf(os.Args[2], "%d", &n)
	if err != nil || n < 0 {
		fmt.Fprintln(os.Stderr, "invalid count")
		os.Exit(1)
	}

	files := os.Args[3:]
	exitCode := 0

	// Print blank line before a successful header if anything was printed before
	printedSomething := false

	for _, fname := range files {
		data, err := os.ReadFile(fname)
		if err != nil {
			// Print the error
			fmt.Println(err)
			exitCode = 1
			printedSomething = true
			continue
		}

		// If we've printed anything before (error or previous file), print blank line
		if printedSomething {
			fmt.Println()
		}
		printedSomething = true

		// Header for successful files
		fmt.Printf("==> %s <==\n", fname)

		// Tail -c n
		if n < len(data) {
			data = data[len(data)-n:]
		}

		fmt.Print(string(data))
	}

	os.Exit(exitCode)
}
