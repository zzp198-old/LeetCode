package main

func restoreMatrix(rowSum []int, colSum []int) [][]int {
	ans := make([][]int, len(rowSum))
	for i, rs := range rowSum {
		ans[i] = make([]int, len(colSum)) // 逐行选min,选完一个后横竖对应的和减去
		for j, cs := range colSum {
			ans[i][j] = min(rs, cs)
			rs -= ans[i][j]
			colSum[j] -= ans[i][j]
		}
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
