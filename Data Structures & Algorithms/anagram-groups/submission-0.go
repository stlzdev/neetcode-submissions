func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return [][]string{}
	}
	out := make([][]string, 0)
	groups := make(map[string][]string)
	for _, str1 := range(strs) {
		key := sortString(str1)
		groups[key] = append(groups[key], str1)
	}
	for _, group := range groups {
		out = append(out, group)
	}
	return out
}

func sortString(s string) string {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		for j := i + 1; j < len(runes); j++ {
			if runes[i] > runes[j] {
				runes[i], runes[j] = runes[j], runes[i]
			}
		}
	}
	return string(runes)
}