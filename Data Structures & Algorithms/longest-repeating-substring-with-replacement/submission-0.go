func characterReplacement(s string, k int) int {
	maxlen := 0
	fmap := make(map[byte]int) 
	lpt := 0
	maxFreq := 0
	size := 0
	for rpt := range len(s) {
		size = rpt - lpt + 1
		fmap[s[rpt]]++
		maxFreq = max(maxFreq, fmap[s[rpt]])
		for size - maxFreq > k {
			fmap[s[lpt]]--
			lpt++
			maxFreq = max(maxFreq, fmap[s[rpt]])
			size = rpt - lpt + 1
		}
		maxlen = max(maxlen, size)
	}
	return maxlen
}
