func countSubstrings(s string) int {
    if len(s) == 1 {
		return 1
	}
	count := 0
	var expand func(int, int)
	expand = func(start int, end int) {
		for start >= 0 && end < len(s) && s[start] == s[end] {
			count += 1
			start -= 1
			end += 1
		}
	}
	for i := range len(s) {
		expand(i, i)
		expand(i, i+1)
	}
	return count
}
