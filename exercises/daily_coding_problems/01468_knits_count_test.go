/*
* A knight's tour is a sequence of moves by a knight on a chessboard such that all squares are visited once.
*
* Given N, write a function to return the number of knight's tours on an N by N chessboard.
* */

package dailycodingproblems

import (
	"testing"
)

var (
	xMove = [8]int{2, 1, -1, -2, -2, -1, 1, 2}
	yMove = [8]int{1, 2, 2, 1, -1, -2, -2, -1}
)

func isSafe(x, y, N int, board [][]int) bool {
	return x >= 0 && y >= 0 && x < N && y < N && board[x][y] == -1
}

func solveKTUtil(x, y, movei, N int, board [][]int) int {
	if movei == N*N {
		return 1
	}

	count := 0
	for k := 0; k < 8; k++ {
		nextX := x + xMove[k]
		nextY := y + yMove[k]
		if isSafe(nextX, nextY, N, board) {
			board[nextX][nextY] = movei
			count += solveKTUtil(nextX, nextY, movei+1, N, board)
			board[nextX][nextY] = -1
		}
	}
	return count
}

func countKnightsTours(N int) int {
	board := make([][]int, N)
	for i := range board {
		board[i] = make([]int, N)
		for j := range board[i] {
			board[i][j] = -1
		}
	}

	board[0][0] = 0
	return solveKTUtil(0, 0, 1, N, board)
}

func TestCountKnightsTours(t *testing.T) {
	N := 5
	t.Logf("Number of knight's tours on a %dx%d board: %d\n", N, N, countKnightsTours(N))
}
