package main

import (
	"fmt"

	search "github.com/Evergreenies/go-algorithms/algos/searching"
)

func main() {
	arr := []int{-4, -2, -7, -1, -9, -5, -3}
	target := -9
	fmt.Println(search.LinearSearch(arr, target))
}
