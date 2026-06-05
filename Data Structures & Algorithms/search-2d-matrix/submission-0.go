func searchMatrix(matrix [][]int, target int) bool {
	h := len(matrix)
	w := len(matrix[0])
	l := 0
	r := h * w - 1
	for l <= r {
		m := l + (r - l) / 2
		row, col := mapidxtopos(w, m)
		el := matrix[row][col]
		if el == target {
			return true
		} else if el < target {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return false
}

func mapidxtopos(w int, idx int) (int, int) {
	row := idx / w
	col := idx % w
	return row, col
}