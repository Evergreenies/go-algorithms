/**
* You have a large array with most of the elements as zero.
*
* Use a more space-efficient data structure, SparseArray, that implements the same interface:
*
* init(arr, size): initialize with the original large array and size.
* set(i, val): updates index at i with val.
* get(i): gets the value at index i.
* */

package dailycodingproblems

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type sparseArray struct {
	arr  map[int]int
	size int
}

func (sa *sparseArray) init(arr []int, size int) {
	sa.arr = make(map[int]int)
	sa.size = size

	for index, value := range arr {
		if value != 0 {
			sa.arr[index] = value
		}
	}
}

func (sa *sparseArray) set(index int, value int) {
	if index < 0 || index >= sa.size {
		fmt.Println("index out of range")

		return
	}

	if value == 0 {
		delete(sa.arr, index)
	} else {
		sa.arr[index] = value
	}
}

func (sa *sparseArray) get(index int) (int, error) {
	if index < 0 || index >= sa.size {
		fmt.Println("index out of range")

		return -1, errors.New("index out of range")
	}

	return sa.arr[index], nil
}

func TestSparseArray(t *testing.T) {
	assert := assert.New(t)
	arr := []int{0, 0, 3, 0, 0, 4, 0, 5}
	size := len(arr)

	sparseArr := new(sparseArray)
	sparseArr.init(arr, size)

	item, err := sparseArr.get(5)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(4, item, fmt.Sprintf("expected 4; got %v\n", item))

	_, err = sparseArr.get(4)
	if err != nil {
		t.Log(err)
	}

	sparseArr.set(4, 12)
	item, err = sparseArr.get(4)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(12, item, fmt.Sprintf("expected 12; got %v\n", item))
}
