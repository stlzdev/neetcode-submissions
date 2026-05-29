func lengthOfLongestSubstring(s string) int {
	curr := 0 
	hmap := make(map[byte]struct{})
	lpt := 0
	for rpt := range len(s) {
		for {
			_, exist := hmap[s[rpt]]
			if !exist {
				break
			}
			delete(hmap, s[lpt])
			lpt++
		}
		hmap[s[rpt]] = struct{}{}
		curr = max(curr, len(hmap))
	}
	return curr
}