package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 10 {
		fmt.Println("Error")
		return
	}

	board := make([][]rune, 9)

	for a := 1; a < 10; a++ {
		args := os.Args[a]

		if len(args) != 9 {
			fmt.Println("Error")
			return
		}

		row := make([]rune, 9)

		for j, v := range args {

			if v != '.' && (v < '1' || v > '9') {
				fmt.Println("Error")
				return
			}

			row[j] = v
		}

		board[a-1] = row
	}

	if !isValidBoard(board) {
		fmt.Println("Error")
		return
	}

	if solveBoard(board) {
		for i := 0; i < len(board); i++ {
			for j := 0; j < len(board[i]); j++ {
				fmt.Printf("%c", board[i][j])
				if j != len(board)-1 {
					fmt.Print(" ")
				}
			}
			fmt.Println()
		}
	} else {
		fmt.Println("Error")
	}
}

func getNextEmpty(sudoku [][]rune) (int, int, error) {
	for j := 0; j < 9; j++ {
		for i := 0; i < 9; i++ {

			if sudoku[j][i] == '.' {
				return j, i, nil
			}
		}
	}

	return 0, 0, errors.New("Error")
}

func isValidBoard(sudoku [][]rune) bool {

	for i := 0; i < 9; i++ {

		rowSet := make(map[rune]bool)
		colSet := make(map[rune]bool)
		for j := 0; j < 9; j++ {
			if sudoku[i][j] != '.' {
				if rowSet[sudoku[i][j]] {
					return false
				}

				rowSet[sudoku[i][j]] = true
			}
			if sudoku[j][i] != '.' {

				if colSet[sudoku[j][i]] {
					return false
				}

				colSet[sudoku[j][i]] = true
			}
		}
	}

	for i := 0; i < 9; i += 3 {

		for j := 0; j < 9; j += 3 {
			gridSet := make(map[rune]bool)
			for x := i; x < i+3; x++ {
				for y := j; y < j+3; y++ {
					if sudoku[x][y] != '.' {
						if gridSet[sudoku[x][y]] {
							return false
						}
						gridSet[sudoku[x][y]] = true
					}
				}
			}
		}
	}

	return true
}

func solveBoard(board [][]rune) bool {
	j, i, e := getNextEmpty(board)
	if e != nil {
		return true
	}

	for k := '1'; k <= '9'; k++ {

		board[j][i] = k

		if isValidBoard(board) && solveBoard(board) {
			return true
		}

		board[j][i] = '.'
	}

	return false
}