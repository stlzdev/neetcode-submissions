func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	vals := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		vals[s[i]]++
		vals[t[i]]--
	}
	for _, v := range vals {
		if v != 0 {
			return false
		}
	}
	return true
}
