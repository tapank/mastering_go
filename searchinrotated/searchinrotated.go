package searchinrotated

import "fmt"

func Search(nums []int, target int) int {
	return search(nums, target)
}

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	length, offset := len(nums), findmin(nums)
	fmt.Println("min:", findmin(nums))
	l, r := 0, length-1
	for l <= r {
		m := (l + r) / 2
		mval := nums[tr(m, offset, length)]
		if target > mval {
			l = m + 1
		} else if target < mval {
			r = m - 1
		} else {
			return tr(m, offset, length)
		}
	}
	return -1
}

func tr(index, offset, len int) int {
	return (index + offset) % len
}

func findmin(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	lastindex := len(nums) - 1
	if nums[0] < nums[lastindex] {
		return 0
	}
	l, r := 0, lastindex
	for l <= r {
		m := (l + r) / 2
		if nums[l] < nums[m] {
			l = m
		} else if nums[m] < nums[r] {
			r = m
		} else {
			return m + 1
		}
	}
	return -1
}
