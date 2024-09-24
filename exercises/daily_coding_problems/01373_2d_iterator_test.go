/**
* Implement a 2D iterator class. It will be initialized with an array of arrays, and should
* implement the following methods:
*
* next(): returns the next element in the array of arrays. If there are no more elements, raise an exception.
* has_next(): returns whether or not the iterator still has elements left.
*
* For example, given the input [[1, 2], [3], [], [4, 5, 6]], calling next() repeatedly
* should output 1, 2, 3, 4, 5, 6.
* Do not use flatten or otherwise clone the arrays. Some of the arrays can be empty
* */
package dailycodingproblems

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type iterator2D struct {
	arr      [][]int
	row, col int
}

func newIterator(arr [][]int) *iterator2D {
	return &iterator2D{
		arr: arr,
		row: 0,
		col: 0,
	}
}

func (it *iterator2D) has_next() bool {
	for it.row < len(it.arr) && it.col >= len(it.arr[it.row]) {
		it.row++
		it.col = 0
	}

	return it.row < len(it.arr)
}

func (it *iterator2D) next() (int, error) {
	if !it.has_next() {
		return 0, errors.New("no more element")
	}

	val := it.arr[it.row][it.col]
	it.col++

	return val, nil
}

func TestIterator2D(t *testing.T) {
	assert := assert.New(t)

	data := [][]int{
		{1, 2},
		{3},
		{},
		{4, 5, 6},
	}

	iterator := newIterator(data)

	for iterator.has_next() {
		value, err := iterator.next()
		if err == nil {
			fmt.Println(value)
		} else {
			assert.NotNil(err)
		}
	}
}
