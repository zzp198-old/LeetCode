package problems

// 除了模拟遍历外,考察数根,即各位一直相加后的一位数
func addDigits(num int) int {
	return (num-1)%9 + 1
}
