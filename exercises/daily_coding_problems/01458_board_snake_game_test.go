/*
* Snakes and Ladders is a game played on a 10 x 10 board, the goal of which is get from
* square 1 to square 100. On each turn players will roll a six-sided die and move forward
* a number of spaces equal to the result. If they land on a square that represents a snake
* or ladder, they will be transported ahead or behind, respectively, to a new square.
*
* Find the smallest number of turns it takes to play snakes and ladders.
*
* For convenience, here are the squares representing snakes and ladders, and their outcomes:
*
* snakes = {16: 6, 48: 26, 49: 11, 56: 53, 62: 19, 64: 60, 87: 24, 93: 73, 95: 75, 98: 78}
* ladders = {1: 38, 4: 14, 9: 31, 21: 42, 28: 84, 36: 44, 51: 67, 71: 91, 80: 100}
* */
package dailycodingproblems

import "testing"

type snakeBoardStrct struct{}

func (snakeBoardStrct) minMoves(snakes, leaders map[int]int) int {
	type QueueNode struct {
		pos, moves int
	}

	board := make([]int, 101)
	for index := 1; index <= 100; index++ {
		board[index] = index
	}

	for start, end := range snakes {
		board[start] = end
	}
	for start, end := range leaders {
		board[start] = end
	}

	visited := make([]bool, 101)
	queue := []QueueNode{
		{1, 0},
	}
	visited[1] = true

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node.pos == 100 {
			return node.moves
		}

		for dice := 1; dice <= 6; dice++ {
			newPos := node.pos + dice

			if newPos <= 100 && !visited[board[newPos]] {
				visited[board[newPos]] = true
				queue = append(queue, QueueNode{board[newPos], node.moves + 1})
			}
		}
	}

	return -1
}

func TestSnakeBoard(t *testing.T) {
	snakes := map[int]int{16: 6, 48: 26, 49: 11, 56: 53, 62: 19, 64: 60, 87: 24, 93: 73, 95: 75, 98: 78}
	ladders := map[int]int{1: 38, 4: 14, 9: 31, 21: 42, 28: 84, 36: 44, 51: 67, 71: 91, 80: 100}

	nb := snakeBoardStrct{}
	t.Log(nb.minMoves(snakes, ladders))
}
