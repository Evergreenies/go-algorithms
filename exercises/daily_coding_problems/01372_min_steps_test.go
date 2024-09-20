/*
You are given an M by N matrix consisting of booleans that represents a board.
Each True boolean represents a wall. Each False boolean represents a tile you can walk on.

Given this matrix, a start coordinate, and an end coordinate, return the minimum number of
steps required to reach the end coordinate from the start. If there is no possible path,
then return null. You can move up, left, down, and right. You cannot move through walls.
You cannot wrap around the edges of the board.

For example, given the following board:

[[f, f, f, f],
[t, t, f, t],
[f, f, f, f],
[f, f, f, f]]
and start = (3, 0) (bottom left) and end = (0, 0) (top left), the minimum number of steps
required to reach the end is 7, since we would need to go through (1, 2) because there is
a wall everywhere else on the second row.
* */

package dailycodingproblems

import (
	"container/list"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type minStepsReq struct{}

type point struct {
	row, col, steps int
}

func (minStepsReq) minSteps(board [][]bool, start, end point) int {
	rows := len(board)
	if rows < 0 {
		return -1
	}

	cols := len(board[0])
	if cols < 0 {
		return -1
	}

	directions := []point{
		{-1, 0, 0}, // up
		{1, 0, 0},  // down
		{0, -1, 0}, // left
		{0, 1, 0},  // right
	}

	// using queue as we are following BFS also
	queue := list.New()
	queue.PushBack(point{start.row, start.col, 0})

	// memoization
	visited := make([][]bool, rows)
	for index := range visited {
		visited[index] = make([]bool, cols)
	}
	visited[start.row][start.col] = true

	// BFS
	for queue.Len() > 0 {
		element := queue.Front()
		queue.Remove(element)
		current := element.Value.(point)

		if current.row == end.row && current.col == end.col {
			return current.steps
		}

		// explore all four directions
		for _, d := range directions {
			r, c := current.row+d.row, current.col+d.col

			if r >= 0 && r < rows && c >= 0 && c < cols && !board[r][c] && !visited[r][c] {
				queue.PushBack(point{r, c, current.steps + 1})
				visited[r][c] = true
			}
		}
	}

	return -1
}

func TestMinSteps(t *testing.T) {
	assert := assert.New(t)

	board := [][]bool{
		{false, false, false, false},
		{true, true, false, true},
		{false, false, false, false},
		{false, false, false, false},
	}

	start := point{3, 0, 0} // bottom left
	end := point{0, 0, 0}
	mn := minStepsReq{}
	steps := mn.minSteps(board, start, end)

	assert.Equal(steps, 7, fmt.Sprintf("expected 7; got %v\n", steps))
}
