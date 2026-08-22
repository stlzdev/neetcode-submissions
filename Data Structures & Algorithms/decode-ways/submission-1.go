func numDecodings(s string) int {
	dp := make(map[int]int)
	var parse func(int) int
	parse = func(idx int) int {
		if idx == len(s) {
			return 1
		}
		if val, ok := dp[idx]; ok {
			return val
		}
		single := 0
		if s[idx] != '0' {
			single = parse(idx+1)
		}
		double := 0
		if idx+1 < len(s) {
			two, _ := strconv.Atoi(s[idx:idx+2])
			if two >= 10 && two <= 26 {
				double = parse(idx+2)
			}
		}
		dp[idx] = single + double
		return dp[idx]
	}
	return parse(0)
}
