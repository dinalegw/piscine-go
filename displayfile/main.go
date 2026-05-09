package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:] // skip the program name

	if len(args) == 0 {
		fmt.Println("File name missing")
		return
	} else if len(args) > 1 {
		fmt.Println("Too many arguments")
		return
	}

	filename := args[0]

	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Print(string(data))
}
