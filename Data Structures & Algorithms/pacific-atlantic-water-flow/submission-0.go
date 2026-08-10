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
	q := [][2]int{}
	// add all pacific coasts 
	for c := range cnum {
		q = append(q, [2]int{0, c})
		pacific[0][c] = true
	}
	if rnum >= 1 {
		for r := 1; r < rnum; r++ {
			q = append(q, [2]int{r, 0})
			pacific[r][0] = true
		}
	}
	// record all cells reachable from pacific
	for len(q) > 0 {
		head := q[0]
		q = q[1:]
		r, c := head[0], head[1]
		for _, dir := range dirs {
			nr, nc := r + dir[0], c + dir[1]
			if nr < 0 || nr >= rnum || nc < 0 || nc >= cnum || heights[nr][nc] < heights[r][c] || pacific[nr][nc] {
				continue
			}
			q = append(q, [2]int{nr, nc})
			pacific[nr][nc] = true
		}
	}
	// add all atlantic coasts 
	for c := range cnum {
		q = append(q, [2]int{rnum-1, c})
		atlantic[rnum-1][c] = true
	}
	for r := 0; r < rnum-1; r++ {
		q = append(q, [2]int{r, cnum-1})
		atlantic[r][cnum-1] = true
	}
	// record all cells reachable from atlantic
	for len(q) > 0 {
		head := q[0]
		q = q[1:]
		r, c := head[0], head[1]
		for _, dir := range dirs {
			nr, nc := r + dir[0], c + dir[1]
			if nr < 0 || nr >= rnum || nc < 0 || nc >= cnum || heights[nr][nc] < heights[r][c] || atlantic[nr][nc] {
				continue
			}
			q = append(q, [2]int{nr, nc})
			atlantic[nr][nc] = true
		}
	}
	// append only cells reachable from both (status[r][c] == 2)
	for r := range rnum {
		for c := range cnum {
			if pacific[r][c] && atlantic[r][c] {
				out = append(out, []int{r, c})
			}
		}
	}
	return out
}
