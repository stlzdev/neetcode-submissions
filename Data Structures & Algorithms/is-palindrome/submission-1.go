func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	lpt := 0
	rpt := len(s) - 1
	for {
		if lpt + 1 > rpt {
			break
		}
		l := rune(s[lpt]); r := rune(s[rpt])
		if !unicode.IsLetter(l) && !unicode.IsDigit(l) {
			lpt++
			continue
		} 
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			rpt--
			continue
		} 
		if l == r {
			lpt++
			rpt--
		} else {
			return false
		}
	}
	return true
}