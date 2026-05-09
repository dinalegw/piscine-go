package piscine

import "github.com/01-edu/z01"

func EightQueens() {
	var board [8]int
	solve(board, 0)
}

// Recursive solver
func solve(board [8]int, col int) {
	if col == 8 {
		printBoard(board)
		return
	}
	for row := 0; row < 8; row++ {
		board[col] = row
		if isSafe(board, col) {
			solve(board, col+1)
		}
	}
}

// Check if the queen placement is safe
func isSafe(board [8]int, col int) bool {
	for i := 0; i < col; i++ {
		if board[i] == board[col] || abs(board[i]-board[col]) == col-i {
			return false
		}
	}
	return true
}

// Print a solution in the required format
func printBoard(board [8]int) {
	for i := 0; i < 8; i++ {
		z01.PrintRune(rune(board[i] + 1 + '0'))
	}
	z01.PrintRune('\n')
}

// Absolute value helper
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
