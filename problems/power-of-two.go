package problems

// 2的幂  1开头,后跟n个0,0次幂特殊      10 100 1000 10000
func isPowerOfTwo(n int) bool {
	return n > 0 && n&(n-1) == 0
}
