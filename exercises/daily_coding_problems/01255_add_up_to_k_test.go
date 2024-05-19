package dailycodingproblems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
*  Given a list of numbers and a number k, return whether any two numbers from the list add up to k.
*
* For example, given [10, 15, 3, 7] and k of 17, return true since 10 + 7 is 17.
* **/

func add_up_to_k(arr []int, target int) bool {
	if len(arr) == 0 {
		return false
	}

	visited := make(map[int]bool)
	for _, item := range arr {
		sub := target - item
		visited[item] = true

		if sub == 0 {
			return true
		}

		if _, ok := visited[sub]; ok {
			return true
		}

		visited[sub] = true
	}

	return false
}

func TestAddUpToK(t *testing.T) {
	assert := assert.New(t)
	assert.True(add_up_to_k([]int{10, 15, 3, 7}, 17), "it must add up to K")
	assert.True(add_up_to_k([]int{10, 15, 3, 7}, 18), "it must add up to K")
	assert.False(add_up_to_k([]int{10, 15, 3, 7}, 19), "it may not add up to K")
}
