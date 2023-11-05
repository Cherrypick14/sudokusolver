package main

import (
	"github.com/01-edu/z01"
	"os"
)

// Define the size of the Sudoku grid (typically 9x9)
const N = 9

func main() {
	if len(os.Args) != N+1 {
		z01.PrintRune('E')
		z01.PrintRune('r')
		z01.PrintRune('r')
		z01.PrintRune('o')
		z01.PrintRune('r')
		z01.PrintRune('\n')
		return
	}

	board := make([][]int, N)
	for i := range board {
		board[i] = make([]int, N)
	}

	for i := 1; i <= N; i++ {
		if len(os.Args[i]) != N {
			z01.PrintRune('E')
			z01.PrintRune('r')
			z01.PrintRune('r')
			z01.PrintRune('o')
			z01.PrintRune('r')
			z01.PrintRune('\n')
			return
		}
		for j := 0; j < N; j++ {
			if os.Args[i][j] == '.' {
				board[i-1][j] = 0
			} else if os.Args[i][j] >= '1' && os.Args[i][j] <= '9' {
				board[i-1][j] = int(os.Args[i][j] - '0')
			} else {
				z01.PrintRune('E')
				z01.PrintRune('r')
				z01.PrintRune('r')
				z01.PrintRune('o')
				z01.PrintRune('r')
				z01.PrintRune('\n')

				return
			}
		}
	}

	if solveSudoku(board) {
		printBoard(board)
	} else {
		z01.PrintRune('E')
		z01.PrintRune('r')
		z01.PrintRune('r')
		z01.PrintRune('o')
		z01.PrintRune('r')
		z01.PrintRune('\n')

	}
}

func solveSudoku(board [][]int) bool {
	var row, col int
	empty := true

	// Find an empty cell
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if board[i][j] == 0 {
				row, col = i, j
				empty = false
				break
			}
		}
		if !empty {
			break
		}
	}

	// No empty cell found, the puzzle is solved
	if empty {
		return true
	}

	// Try placing a number (1 to 9) in the empty cell
	for num := 1; num <= 9; num++ {
		if isSafe(board, row, col, num) {
			board[row][col] = num

			// Recursively try to solve the rest of the puzzle
			if solveSudoku(board) {
				return true
			}

			// If the number doesn't lead to a solution, backtrack
			board[row][col] = 0
		}
	}

	// No valid number found, return false to trigger backtracking
	return false
}

func isSafe(board [][]int, row, col, num int) bool {
	return !usedInRow(board, row, num) &&
		!usedInCol(board, col, num) &&
		!usedInBox(board, row-row%3, col-col%3, num)
}

func usedInRow(board [][]int, row, num int) bool {
	for col := 0; col < N; col++ {
		if board[row][col] == num {
			return true
		}
	}
	return false
}

func usedInCol(board [][]int, col, num int) bool {
	for row := 0; row < N; row++ {
		if board[row][col] == num {
			return true
		}
	}
	return false
}

func usedInBox(board [][]int, startRow, startCol, num int) bool {
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			if board[row+startRow][col+startCol] == num {
				return true
			}
		}
	}
	return false
}

func printBoard(board [][]int) {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			z01.PrintRune(rune(board[i][j] + '0'))
			if j < N-1 {
				z01.PrintRune(' ')

			}
		}
		z01.PrintRune('\n')

	}
	z01.PrintRune('\n')
}