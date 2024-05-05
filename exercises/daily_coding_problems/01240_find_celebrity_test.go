package dailycodingproblems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
* At a party, there is a single person who everyone knows, but who does not know anyone
* in return (the "celebrity"). To help figure out who this is, you have access to an O(1)
* method called knows(a, b), which returns True if person a knows person b, else False.
*
* Given a list of N people and the above operation, find a way to identify the celebrity
* in O(N) time.
* **/

func findCelebrity(people [][]int) int {
	length := len(people)
	candidate := 0

	for index := 1; index < length; index++ {
		if people[index][candidate] == 0 {
			candidate = index
		}
	}

	for _, value := range people[candidate] {
		if value == 1 {
			return -1
		}
	}

	return candidate
}

func TestFindCelebrity(t *testing.T) {
	assert := assert.New(t)

	people := [][]int{
		{0, 0, 1, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 0},
		{0, 0, 1, 0},
	}

	assert.Equal(findCelebrity(people), 2, "here we have 2nd is celebrity person")

	people = [][]int{
		{0, 0, 1, 0},
		{0, 0, 1, 0},
		{0, 0, 1, 0},
		{0, 0, 1, 0},
		{0, 0, 1, 0},
	}

	assert.Equal(findCelebrity(people), -1, "there is no one celebrity in party today")
}
