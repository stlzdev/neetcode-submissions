func islandsAndTreasure(grid [][]int) {
	q := [][]int{}
	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for r := range len(grid) {
		for c := range len(grid[0]) {
			if grid[r][c] == 0 {
				q = append(q, []int{r, c})
			}
		}
	}
	for len(q) > 0 {
		head := q[0]
		r := head[0]; c := head[1]
		q = q[1:]
		for _, dir := range dirs {
			nr, nc := r+dir[0], c+dir[1]
			if nr < 0 || nc < 0 || nr >= len(grid) || nc >= len(grid[0]) || grid[nr][nc] != 2147483647 {
				continue
			}
			grid[nr][nc] = grid[r][c] + 1
			q = append(q, []int{nr, nc})
		}
	}
}