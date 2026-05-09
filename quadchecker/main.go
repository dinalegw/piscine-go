package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

// Quad defines a pattern with specific corner and border characters
type Quad struct {
	Name     string
	TopLeft  rune
	TopRight rune
	BotLeft  rune
	BotRight rune
	Horiz    rune
	Vert     rune
}

// All possible quad patterns based on your quad functions
var quads = []Quad{
	{"quadA", 'o', 'o', 'o', 'o', '-', '|'},
	{"quadB", '/', '\\', '\\', '/', '*', '*'},
	{"quadC", 'A', 'A', 'C', 'C', 'B', 'B'},
	{"quadD", 'A', 'C', 'A', 'C', 'B', 'B'},
	{"quadE", 'A', 'C', 'C', 'A', 'B', 'B'},
}

// readInput reads all lines from stdin
func readInput() ([]string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

// isValidRectangle checks if all lines have the same length and contain only printable characters
func isValidRectangle(lines []string) bool {
	if len(lines) == 0 {
		return false
	}

	width := len(lines[0])
	if width == 0 {
		return false
	}

	// Check all lines have same length
	for _, line := range lines {
		if len(line) != width {
			return false
		}
		// Check for non-printable characters (except space)
		for _, ch := range line {
			if ch < 32 || ch > 126 {
				if ch != ' ' && ch != '\t' {
					return false
				}
			}
		}
	}

	return true
}

// checkQuad checks if the input matches a specific quad pattern
func checkQuad(lines []string, quad Quad) bool {
	height := len(lines)
	width := len(lines[0])

	// Special case: 1x1 rectangle
	if width == 1 && height == 1 {
		return lines[0][0] == byte(quad.TopLeft)
	}

	// Check all four corners
	if rune(lines[0][0]) != quad.TopLeft {
		return false
	}
	if rune(lines[0][width-1]) != quad.TopRight {
		return false
	}
	if rune(lines[height-1][0]) != quad.BotLeft {
		return false
	}
	if rune(lines[height-1][width-1]) != quad.BotRight {
		return false
	}

	// Check top and bottom borders (excluding corners)
	for col := 1; col < width-1; col++ {
		if rune(lines[0][col]) != quad.Horiz || rune(lines[height-1][col]) != quad.Horiz {
			return false
		}
	}

	// Check left and right borders (excluding corners)
	for row := 1; row < height-1; row++ {
		if rune(lines[row][0]) != quad.Vert || rune(lines[row][width-1]) != quad.Vert {
			return false
		}
	}

	// Check interior is only spaces
	for row := 1; row < height-1; row++ {
		for col := 1; col < width-1; col++ {
			if lines[row][col] != ' ' {
				return false
			}
		}
	}

	return true
}

// findMatchingQuads returns all quads that match the input pattern
func findMatchingQuads(lines []string) []string {
	var matches []string
	height := len(lines)
	width := len(lines[0])

	for _, quad := range quads {
		if checkQuad(lines, quad) {
			matches = append(matches, fmt.Sprintf("[%s] [%d] [%d]", quad.Name, width, height))
		}
	}

	return matches
}

func main() {
	// Read input from stdin
	lines, err := readInput()
	if err != nil || len(lines) == 0 {
		fmt.Println("Not a quad function")
		return
	}

	// Validate input is a proper rectangle
	if !isValidRectangle(lines) {
		fmt.Println("Not a quad function")
		return
	}

	// Find all matching quads
	matches := findMatchingQuads(lines)

	// Handle output based on matches
	if len(matches) == 0 {
		fmt.Println("Not a quad function")
	} else {
		// Sort matches alphabetically
		sort.Strings(matches)
		// Join with " || " separator
		result := strings.Join(matches, " || ")
		fmt.Println(result)
	}
}
