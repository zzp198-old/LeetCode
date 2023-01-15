package problems

func islandPerimeter(grid [][]int) (answer int) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				// 上
				if i == 0 || grid[i-1][j] == 0 {
					answer++
				}
				// 左
				if j == 0 || grid[i][j-1] == 0 {
					answer++
				}
				// 下
				if i == len(grid)-1 || grid[i+1][j] == 0 {
					answer++
				}
				// 右
				if j == len(grid[i])-1 || grid[i][j+1] == 0 {
					answer++
				}
			}
		}
	}
	return
}
