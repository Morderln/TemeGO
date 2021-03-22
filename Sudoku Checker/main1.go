package main

import (
	"fmt"
)


func makeMatrix(board Matrix) {
	lines := [9]int {}
	columns := [9] int {}
	squares := [9] int {}
	for i := range board {
		for j := range board[i] {
			var value = byteNumbers[board[i][j]]
			lines[i] += value
			columns[j] += value
			squares[findSquare(i, j)] += value
		}
	}
	if checkSudoku(lines) {
		fmt.Println("This sudoku is correct")
	} else {
		fmt.Println("This sudoku is incorrect")
	}
	fmt.Println(lines)
	fmt.Println(columns)
	fmt.Println(squares)
}

func checkSudoku(example [9] int) bool{
	for i:=range example {
		if example[i]!=511 {
			return false
		}
	}
	return true
}



func main() {
	makeMatrix(board)
}
