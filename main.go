package main

import (
	"fmt"

	search "github.com/Evergreenies/go-algorithms/algos/searching"
)

func main() {
	arr := []int{-4, -2, -7, -1, -9, -5, -3}
	target := -9
	fmt.Println(search.BinarySearch(arr, target))

	arr = []int{0, 0, 0, 1, 3, 5, 7, 9, 11, 13}
	target = 7
	fmt.Println(search.BinarySearch(arr, target))
}
