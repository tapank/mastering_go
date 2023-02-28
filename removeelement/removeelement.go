package removeelement

func removeElement(nums []int, val int) []int {
	ret := []int{}
	for _, v := range nums {
		if val != v {
			ret = append(ret, v)
		}
	}
	return ret
}
