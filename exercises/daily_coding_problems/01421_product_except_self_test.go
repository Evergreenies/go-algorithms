/**
* Given an array of integers, return a new array such that each element at index i of the new array
* is the product of all the numbers in the original array except the one at i.
*
* For example, if our input was [1, 2, 3, 4, 5], the expected output would be [120, 60, 40, 30, 24].
* If our input was [3, 2, 1], the expected output would be [2, 3, 6].
*
* Follow-up: what if you can't use division?
* */
package dailycodingproblems

import "testing"

type productExceptSelfStrct struct{}

func (productExceptSelfStrct) productExceptSelf(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}

	leftProduct, rightProduct, output := make([]int, length), make([]int, length), make([]int, length)

	leftProduct[0] = 1
	for index := 1; index < length; index++ {
		leftProduct[index] = leftProduct[index-1] * arr[index-1]
	}

	rightProduct[length-1] = 1
	for index := length - 2; index >= 0; index-- {
		rightProduct[index] = rightProduct[index+1] * arr[index+1]
	}

	for index := 0; index < length; index++ {
		output[index] = leftProduct[index] * rightProduct[index]
	}

	return output
}

func TestProductExceptSelf(t *testing.T) {
	p := productExceptSelfStrct{}
	arr := p.productExceptSelf([]int{1, 2, 3, 4, 5})
	t.Logf("got %v\n", arr)
}
