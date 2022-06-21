package maxarea

func MaxArea(height []int) int {
	max := 0
	for i := 0; i < len(height); i++ {
		if height[i] == 0 {
			continue
		}
		for j := len(height) - 1; j > i; j-- {
			if height[j] == 0 {
				continue
			}
			len := j - i
			ht := height[i]
			if height[j] < ht {
				ht = height[j]
			}
			area := len * ht
			if area > max {
				max = area
			} else if height[i] <= height[j] {
				break
			}
		}
	}
	return max
}
