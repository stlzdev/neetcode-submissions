func maxProduct(nums []int) int {
    if len(nums) == 0 {
		return 0
	}
	gloMax := nums[0]
	currMax := nums[0]
	currMin := nums[0]
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if num < 0 {
			currMax, currMin = currMin, currMax
		}
		currMax = max(num, currMax * num)
		currMin = min(num, currMin * num)
		gloMax = max(gloMax, currMax)
	}
	return gloMax
}