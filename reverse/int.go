package reverse

const NEGATIVE int32 = -1
const POSITIVE int32 = 1

func Reverse(x int) int {
	var r int32
	if x == 0 {
		return 0
	}
	sign := POSITIVE
	if x < 0 {
		sign = NEGATIVE
		x *= int(sign)
	}
	for x > 0 {
		digit := int32(x % 10)
		savedr := r
		r *= 10
		if r < 0 || r/10 != savedr {
			return 0
		}
		x /= 10
		r += digit
	}
	r *= sign
	return int(r)
}
