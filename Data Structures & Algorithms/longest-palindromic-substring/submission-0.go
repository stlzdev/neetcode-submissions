func longestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}
	start, end := 0, 0
	var expand func(int, int) int
	expand = func(start int, end int) int {
		for start >= 0 && end < len(s) && s[start] == s[end] {
			start -= 1
			end += 1
		}
		return end - start - 1
	}
	for i := range len(s) {
		odd := expand(i, i)
		even := expand(i, i+1)
		max := max(odd, even)
		if max > (end - start) {
			start = i - (max-1)/2
			end = i + max/2
		}
	}
	return s[start:end+1]
}
