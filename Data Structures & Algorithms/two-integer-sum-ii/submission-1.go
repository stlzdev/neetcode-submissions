func twoSum(numbers []int, target int) []int {
	lpt := 0
	rpt := len(numbers) - 1
	for lpt < rpt {
		curr := numbers[lpt] + numbers[rpt]
		if curr == target {
			return []int{lpt + 1, rpt + 1}
		} else if curr < target {
			lpt++
		} else {
			rpt--
		}
	}
	return []int{}
}