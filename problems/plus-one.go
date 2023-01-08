package problems

func plusOne(digits []int) []int {
	p := len(digits) - 1
	digits[p]++
	for ; digits[p] == 10 && p > 0; p-- {
		digits[p] -= 10
		digits[p-1] += 1
	}
	if digits[0] == 10 {
		digits = make([]int, len(digits)+1)
		digits[0] = 1
	}
	return digits
}
