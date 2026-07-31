func generateParenthesis(n int) []string {
	out := []string{}
	var track func(int, int, string)
	track = func(open int, closed int, group string) {
		if len(group) == 2*n {
			out = append(out, group)
			return
		}
		if open < n {
			track(open + 1, closed, group + "(")
		}
		if open	> closed {
			track(open, closed + 1, group + ")")
		}	
	} 
	track(0, 0, "")
	return out
}
