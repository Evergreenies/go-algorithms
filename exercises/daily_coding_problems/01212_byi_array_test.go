package dailycodingproblems

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type BitArray struct {
	arr  []int
	size int
}

type BitArrayIntf interface {
	initBitArray(int) *BitArray

	get(int) int
	set(int, int) error

	isValidIndex(int) bool
	isValidValue(int) bool
}

func newBitArray(size int) *BitArray {
	return new(BitArray).initBitArray(size)
}

func (arr *BitArray) initBitArray(size int) *BitArray {
	return &BitArray{
		arr:  make([]int, size),
		size: size,
	}
}

func (bit *BitArray) isValidIndex(index int) bool {
	if index < 0 && index > bit.size {
		return false
	}

	return true
}

func (bit *BitArray) isValidValue(value int) bool {
	if value == 0 || value == 1 {
		return true
	}

	return false
}

func (bit *BitArray) get(index int) int {
	if bit.isValidIndex(index) {
		return bit.arr[index+1]
	}

	return -1
}

func (bit *BitArray) set(index int, value int) error {
	if bit.isValidIndex(index) && bit.isValidValue(value) {
		bit.arr[index+1] = value

		return nil
	}

	return errors.New("invalid index or value")
}

func TestBitArray(t *testing.T) {
	assert := assert.New(t)

	bit := newBitArray(5)
	bit.set(1, 1)
	assert.Equal(bit.get(1), 1, "value must be 1 at position 1")

	err := bit.set(2, 2)
	assert.Error(err, "invalid index or value")
}
