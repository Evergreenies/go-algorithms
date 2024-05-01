package dailycodingproblems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
* A classroom consists of N students, whose friendships can be represented in an adjacency list.
* For example, the following descibes a situation where 0 is friends with 1 and 2, 3 is friends
* with 6, and so on.

{0: [1, 2],
 1: [0, 5],
 2: [0],
 3: [6],
 4: [],
 5: [1],
 6: [3]}

* Each student can be placed in a friend group, which can be defined as the transitive closure
* of that student's friendship relations. In other words, this is the smallest set such that
* no student in the group has any friends outside this group. For the example above, the friend
* groups would be {0, 1, 2, 5}, {3, 6}, {4}.
*
* Given a friendship list such as the one above, determine the number of friend groups in the class.
* **/

func find_friends_groups_dfs(student int, friendship map[int][]int, visited map[int]bool) {
	if visited[student] {
		return
	}

	visited[student] = true
	for _, friend := range friendship[student] {
		find_friends_groups_dfs(friend, friendship, visited)
	}
}

func count_friends_group(friendship map[int][]int) int {
	groups_count := 0
	visited := map[int]bool{}

	for student := range friendship {
		if !visited[student] {
			find_friends_groups_dfs(student, friendship, visited)

			groups_count += 1
		}
	}

	return groups_count
}

func TestCountFriendsGroup(t *testing.T) {
	assert := assert.New(t)
	friendships := map[int][]int{
		0: {1, 2},
		1: {0, 5},
		2: {0},
		3: {6},
		4: {},
		5: {1},
		6: {3},
	}

	assert.Equal(count_friends_group(friendships), 3, "here we have just 3 friends group")
}
