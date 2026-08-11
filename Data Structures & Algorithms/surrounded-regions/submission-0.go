func solve(board [][]byte) {
	rnum, cnum := len(board), len(board[0])
	var dfs func(int, int)
	dfs = func(r int, c int) {
		if r < 0 || r >= rnum || c < 0 || c >= cnum || board[r][c] != 'O' {
			return
		}
		board[r][c] = 'V'
		dfs(r+1, c)
		dfs(r-1, c)
		dfs(r, c+1)
		dfs(r, c-1)
	}
    for c := range cnum {
		dfs(0, c)
		dfs(rnum-1, c)
	}
	for r := range rnum {
		dfs(r, 0)
		dfs(r, cnum-1)
	}
	for r := range rnum {
		for c := range cnum {
			if board[r][c] == 'O' {
				board[r][c] = 'X'
			} else if board[r][c] == 'V' {
				board[r][c] = 'O'
			}
		}
	}
}
