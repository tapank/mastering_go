package myatoi

import (
	"math"
)

func Convert(s string) (n int) {
	sign := 1
	leadingSpace, leadingSign := true, true
loop:
	for _, v := range s {
		switch v {
		case ' ':
			if !leadingSpace {
				break loop
			}
		case '-':
			leadingSpace = false
			if !leadingSign {
				break loop
			}
			sign = -1
			leadingSign = false
		case '+':
			leadingSpace = false
			if !leadingSign {
				break loop
			}
			sign = 1
			leadingSign = false
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			leadingSpace, leadingSign = false, false

			// check overflow
			var testMultiply int
			if testMultiply = n * 10 / 10; testMultiply != n {
				break loop
			}
			testMultiply = n * 10
			testAdd := testMultiply
			if testAdd = testAdd + int(v-'0') - int(v-'0'); testAdd != testMultiply {
				break loop
			}
			n = n*10 + int(v-'0')
		default:
			break loop
		}
		if _, clamped := clamp(n); clamped {
			break loop
		}
	}
	n *= sign
	n, _ = clamp(n)
	return
}

func clamp(i int) (int, bool) {
	min, max := -1*int(math.Pow(2, 31)), int(math.Pow(2, 31))-1
	if i < min {
		return min, true
	} else if i > max {
		return max, true
	}
	return i, false
}
