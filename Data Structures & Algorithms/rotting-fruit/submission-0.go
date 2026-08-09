func orangesRotting(grid [][]int) int {
	q := [][]int{}
	time := 0
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for r := range len(grid) {
		for c := range len(grid[0]) {
			if grid[r][c] == 2 {
				q = append(q, []int{r, c, 0})
			}
		}
	}
	for len(q) > 0 {
		head := q[0]
		q = q[1:]
		r, c, t := head[0], head[1], head[2]
		for _, dir := range dirs {
			nr, nc := r + dir[0], c + dir[1]
			if nr < 0 || nr >= len(grid) || nc < 0 || nc >= len(grid[0]) || grid[nr][nc] != 1 {
				continue
			}
			grid[nr][nc] = 2
			q = append(q, []int{nr, nc, t+1})
			if time < t+1 {
				time = t+1
			}
		}
	}
    for r := range len(grid) {
		for c := range len(grid[0]) {
			if grid[r][c] == 1 {
				return -1
			}
		}
	}
	return time
}
