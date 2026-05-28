func maxProfit(prices []int) int {
	currMax := 0
	lpt := 0
	for rpt := range len(prices) {
		left := prices[lpt]
		right := prices[rpt]
		if left > right {
			lpt = rpt
		}
		currMax = max(currMax, right - left)
	}
	return currMax
}