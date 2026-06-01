func checkInclusion(s1 string, s2 string) bool {
	var s1freq [26]rune
    for _, c := range s1 {
		s1freq[c - 'a']++
    }
	lpt := 0
	var matchfreq [26]rune
	for rpt, v := range s2 {
		if rpt - lpt + 1 > len(s1) {
			matchfreq[rune(s2[lpt]) - 'a']--
			lpt++
		}
		matchfreq[v - 'a']++
		if matchfreq == s1freq {
			return true
		}
	}
	return false
}