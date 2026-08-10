func pacificAtlantic(heights [][]int) [][]int {
	if len(heights) == 0 || len(heights[0]) == 0 {
		return [][]int{}
	}
	dirs := [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	rnum, cnum := len(heights), len(heights[0])
	pacific := make([][]bool, rnum)
	atlantic := make([][]bool, rnum)
	for i := range rnum {
		pacific[i] = make([]bool, cnum)
		atlantic[i] = make([]bool, cnum)
	}
	out := [][]int{}
	var bfs func([][2]int, [][]bool)
	bfs = func(q [][2]int, ocean [][]bool) {
		for len(q) > 0 {
			head := q[0]
			q = q[1:]
			r, c := head[0], head[1]
			for _, dir := range dirs {
				nr, nc := r + dir[0], c + dir[1]
				if nr < 0 || nr >= rnum || nc < 0 || nc >= cnum || heights[nr][nc] < heights[r][c] || ocean[nr][nc] {
					continue
				}
				q = append(q, [2]int{nr, nc})
				ocean[nr][nc] = true
			}
		}
	}
	q := [][2]int{}
	// add all pacific coasts and record reachable
	for c := range cnum {
		q = append(q, [2]int{0, c})
		pacific[0][c] = true
	}
	for r := 1; r < rnum; r++ {
		q = append(q, [2]int{r, 0})
		pacific[r][0] = true
	}
	bfs(q, pacific)
	// add all atlantic coasts and record reachable
	q = [][2]int{}
	for c := range cnum {
		q = append(q, [2]int{rnum-1, c})
		atlantic[rnum-1][c] = true
	}
	for r := 0; r < rnum-1; r++ {
		q = append(q, [2]int{r, cnum-1})
		atlantic[r][cnum-1] = true
	}
	bfs(q, atlantic)
	// append only cells reachable from both
	for r := range rnum {
		for c := range cnum {
			if pacific[r][c] && atlantic[r][c] {
				out = append(out, []int{r, c})
			}
		}
	}
	return out
}