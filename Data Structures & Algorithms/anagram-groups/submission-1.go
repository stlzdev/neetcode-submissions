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
	chars := []rune(s)
	sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j]
	})
	return string(chars)
}