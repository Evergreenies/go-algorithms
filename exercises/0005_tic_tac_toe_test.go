package exercises

import (
	"fmt"
	"testing"
)

const (
	EMPTY = ' '
	X     = 'X'
	O     = 'O'
)

type Board [3][3]rune

func printBoard(board Board) {
	for index_i := range board {
		for index_j := range board[index_i] {
			fmt.Printf("%c", board[index_i][index_j])
		}
		fmt.Println()
	}
}

func TestTicTacToe(t *testing.T) {
	var board Board = [3][3]rune{
		{EMPTY, EMPTY, EMPTY},
		{EMPTY, EMPTY, EMPTY},
		{EMPTY, EMPTY, EMPTY},
	}

	printBoard(board)
	t.Logf("Ran tic tac toe test\n")
}
