package problems

// 等差数列求和 减去数字,得到空余的数
func missingNumber(nums []int) int {
	n := len(nums)
	sum := n * (n + 1) / 2
	for _, num := range nums {
		sum -= num
	}
	return sum
}
