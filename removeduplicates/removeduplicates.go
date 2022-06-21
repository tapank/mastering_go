package removeduplicates

func Remove(arr []int) int {
	return removeDuplicates(arr)
}

// copy from below

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	count, prev := 1, nums[0]
	for _, n := range nums[1:] {
		if n != prev {
			nums[count] = n
			prev = n
			count++
		}
	}
	return count
}
