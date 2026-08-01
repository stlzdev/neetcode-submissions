func exist(board [][]byte, word string) bool {
	rs := len(board)
	cs := len(board[0])
	wlen := len(word)
	exist := false
	var check func(int, int, int)
	check = func(idx int, r int, c int) {
		if exist {
			return
		}
		cell := board[r][c]
		if idx == wlen - 1 && cell == word[idx] {
			exist = true
			return
		} else if idx == wlen - 1 {
			return
		}
		if cell != word[idx] {
			return
		}
		board[r][c] = '#'
		if r > 0 {
			check(idx + 1, r - 1, c)
		}
		if r < rs - 1 {
			check(idx + 1, r + 1, c)
		}
		if c > 0 {
			check(idx + 1, r, c - 1)
		}
		if c < cs - 1 {
			check(idx + 1, r, c + 1)
		}
		board[r][c] = cell
	} 
	for r := 0; r < rs; r++ {
		for c := 0; c < cs; c++ {
			check(0, r, c)
			if exist {
				return true
			}
		}
	}
	return exist
}
