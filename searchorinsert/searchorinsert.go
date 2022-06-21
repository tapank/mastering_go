package searchorinsert

func SearchOrInsert(nums []int, target int) int {
	return searchInsert(nums, target)
}

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	l, r := 0, len(nums)-1
	for l <= r {
		m := (l + r) / 2
		switch {
		case nums[m] > target:
			r = m - 1
		case nums[m] < target:
			l = m + 1
		default:
			return m
		}
	}
	return r + 1
}
