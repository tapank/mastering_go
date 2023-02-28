package addone

func PlusOne(digits []int) []int {
	return plusOne(digits)
}

func plusOne(digits []int) []int {
	carry := 1
	for i := len(digits) - 1; i >= 0; i-- {
		digit := digits[i] + carry
		digits[i] = digit % 10
		carry = digit / 10
	}
	if carry > 0 {
		digits = append([]int{carry}, digits...)
	}
	return digits
}
