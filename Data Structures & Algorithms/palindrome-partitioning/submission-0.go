func partition(s string) [][]string {
	out := [][]string{}
	curr := []string{}
	var split func(int) 
	split = func(start int) {
		if start == len(s) {
			tmp := make([]string, len(curr))
			copy(tmp, curr)
			out = append(out, tmp)
			return
		}
		for end := start + 1; end <= len(s); end++ {
			if isPalindrome(s[start:end]) {
				curr = append(curr, s[start:end])
				split(end)
				curr = curr[:len(curr) - 1]
			}
		}
	}
	split(0)
	return out
}

func isPalindrome(s string) bool {
	for i, j := 0, len(s) - 1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}