package ispalindrome

import "fmt"

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	} else if x < 10 {
		return true
	} else {
		digits := getDigits(x)
		fmt.Println("digits", digits)
		for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
			if digits[i] != digits[j] {
				return false
			}
		}
	}
	return true
}

func getDigits(x int) (digits []int) {
	for x > 0 {
		digit := x % 10
		x /= 10
		digits = append(digits, digit)
	}
	return
}
