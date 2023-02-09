package problems

func dieSimulator(n int, rollMax []int) int {
	mod := int(1e9 + 7)
	var dfs func(int, int, int) int
	var memo [5001][7][16]int
	rollMax = append([]int{1}, rollMax...)
	for i := 0; i < 5001; i++ {
		for j := 0; j < 7; j++ {
			for k := 0; k < 16; k++ {
				memo[i][j][k] = -1
			}
		}
	}
	dfs = func(i, last, cnt int) int {
		if last != -1 && memo[i][last][cnt] != -1 {
			return memo[i][last][cnt]
		}
		if i == n {
			return 1
		}
		res := 0
		for k := 1; k <= 6; k++ {
			if k == last {
				if cnt+1 < rollMax[k] {
					res = (res + dfs(i+1, k, cnt+1)) % mod
				}
			} else {
				res = (res + dfs(i+1, k, 0)) % mod
			}
		}
		if last != -1 {
			memo[i][last][cnt] = res
		}
		return res
	}
	return dfs(0, -1, 0)
}
