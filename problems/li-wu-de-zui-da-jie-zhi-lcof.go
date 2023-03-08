package problems

func maxValue(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    memo := make([][]int, m)
    for i := range memo {
        memo[i] = make([]int, n)
    }
    var dfs func(int, int) int
    dfs = func(i, j int) int {
        if i < 0 || j < 0 {
            return 0
        }
        p := &memo[i][j]
        if *p > 0 { // grid[i][j] 都是正数，记忆化的 memo[i][j] 必然为正数
            return *p
        }
        *p = max(dfs(i, j-1), dfs(i-1, j)) + grid[i][j]
        return *p
    }
    return dfs(len(grid)-1, len(grid[0])-1)
}

func max(a, b int) int { if a < b { return b }; return a }
