func maxAreaOfIsland(grid [][]int) int {
    if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	maxArea := 0
	var dfs func(int, int) int
	dfs = func(r int, c int) int {
		if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) {
			return 0
		}
		if grid[r][c] == 0 {
			return 0
		}
		grid[r][c] = 0
		return 1 + dfs(r, c-1) + dfs(r, c+1) + dfs(r-1, c) + dfs(r+1, c)
	}
	for r := range len(grid) {
		for c := range len(grid[0]) {
			if grid[r][c] == 1 {
				area := dfs(r, c)
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}
	return maxArea
}
