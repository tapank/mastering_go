package main

import (
	"fmt"
	"masteringgo/searchinrotated"
)

func main() {
	pr := func(arr []int, target int, exp int) {
		got := searchinrotated.Search(arr, target)
		result := "pass"
		if got != exp {
			result = "fail"
		}
		fmt.Printf("Result: %s, input: %v, target: %d, expected: %d, output: %d\n", result, arr, target, exp, got)
	}
	pr([]int{1, 3, 5, 6}, 5, 2)
	pr([]int{6, 1, 3, 5}, 5, 3)
	pr([]int{6, 7, 1, 3, 5}, 5, 4)
	pr([]int{5, 6, 1, 3}, 5, 0)
	pr([]int{5, 6, 7, 1, 3}, 5, 0)
	pr([]int{5, 6, 7, 8, 3}, 5, 0)
	pr([]int{3, 5, 6, 7, 1}, 5, 1)
	pr([]int{1, 3, 5, 6}, 1, 0)
	pr([]int{1, 3, 5, 6}, 6, 3)
	pr([]int{1, 3, 5, 6}, 2, -1)
	pr([]int{1, 3, 5, 6}, 0, -1)
	pr([]int{1, 3, 5, 6}, 7, -1)
	pr([]int{4, 5, 6, 7, 0, 1, 2}, 0, 4)
	pr([]int{4, 5, 6, 7, 8, 0, 1, 2}, 0, 5)
	pr([]int{3, 1}, 1, 1)
	pr([]int{3, 1, 2}, 1, 1)
	pr([]int{5, 1, 2, 3, 4}, 1, 1)
	pr([]int{5, 6, 1, 2, 3, 4}, 1, 2)
}
