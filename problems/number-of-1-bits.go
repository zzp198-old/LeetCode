package problems

// 涉及到一个技巧    A&A-1  会把最后一个位1消掉    00110   00101  ->  00100
func hammingWeight(num uint32) (answer int) {
	for ; num != 0; num = num & (num - 1) {
		answer++
	}
	return
}
