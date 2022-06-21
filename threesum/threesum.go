package threesum

import "fmt"

func ThreeSum(nums []int) [][]int {
	threesums := [][]int{}
	l := len(nums)
	if l < 3 {
		return threesums
	}

	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			for k := j + 1; k < l; k++ {
				if i == j+1 || i == j+k+2 || j == k+1 {
					if nums[i]+nums[j]+nums[k] == 0 {
						fmt.Printf("ignoring!!!! index, value pairs: %d/%d, %d/%d, %d/%d\n", i, nums[i], j, nums[j], k, nums[k])
					}
				} else if nums[i]+nums[j]+nums[k] == 0 {
					fmt.Printf("ignoring!!!! index, value pairs: %d/%d, %d/%d, %d/%d\n", i, nums[i], j, nums[j], k, nums[k])
					threesums = addunique(threesums, sort([]int{nums[i], nums[j], nums[k]}))
				}
			}
		}
	}
	return threesums
}

func addunique(ts [][]int, add []int) [][]int {
	for _, tup := range ts {
		if tup[0] == add[0] && tup[1] == add[1] && tup[2] == add[2] {
			return ts
		}
	}
	return append(ts, add)
}

func sort(tup []int) []int {
	if tup[0] > tup[2] {
		tup[0], tup[2] = tup[2], tup[0]
	}

	if tup[0] > tup[1] {
		tup[0], tup[1] = tup[1], tup[0]
	} else if tup[1] > tup[2] {
		tup[1], tup[2] = tup[2], tup[1]
	}
	return tup
}
