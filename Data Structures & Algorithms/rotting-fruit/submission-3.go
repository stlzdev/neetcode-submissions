func orangesRotting(grid [][]int) int {
	q := [][2]int{}
	time := 0
	dirs := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for r := range len(grid) {
		for c := range len(grid[0]) {
			if grid[r][c] == 2 {
				q = append(q, [2]int{r, c})
			}
		}
	}
	for len(q) > 0 {
		complete := false
		size := len(q)
		for i := 0; i < size; i++ {
			head := q[0]
			q = q[1:]
			r, c := head[0], head[1]
			for _, dir := range dirs {
				nr, nc := r + dir[0], c + dir[1]
				if nr < 0 || nr >= len(grid) || nc < 0 || nc >= len(grid[0]) || grid[nr][nc] != 1 {
					continue
				}
				grid[nr][nc] = 2
				q = append(q, [2]int{nr, nc})
				complete = true
			}
		}
		if complete {
			time += 1
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
