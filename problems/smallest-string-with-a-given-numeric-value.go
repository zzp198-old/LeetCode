package problems

import "strings"

func getSmallestString(n int, k int) string {
	i := n             // i 是开头全为 a 的部分的长度
	for ; i > 0; i-- { // 从最大的长度开始，直到第一个长度 i，使得剩余的 getSmallestString(n-i, k-i) 有解
		if (k-i) >= n-i && k-i <= (n-i)*26 {
			break
		}
	}
	j := n - i         // j 是结尾全为 j 的部分的长度
	for ; j > 0; j-- { // 从最大的长度开始，直到第一个长度 i，使得剩余的中间的 getSmallestString(n-i-j, k-i-26*j) 有解
		if k-i-26*j >= n-i-j && k-i-26*j <= 26*n-i-j {
			break
		}
	}

	leftLength := n - i - j
	leftValue := k - i - 26*j
	middle := ""
	for leftValue > 0 {
		for i := 2; i <= 25; i++ { // 从小字母 b 开始，到 y 结束，找到第一个字母，使得剩余长度及剩余价值的和有解
			if (leftLength-1)*25 >= leftValue-i {
				leftValue -= i
				leftLength--
				middle += string(rune(i - 1 + 'a'))
				break
			}
		}
	}
	return strings.Repeat("a", i) + middle + strings.Repeat("z", j)
}
