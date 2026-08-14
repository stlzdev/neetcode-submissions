func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	nums[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		nums[i] = max(nums[i-2] + nums[i], nums[i-1])
	}
	return nums[len(nums)-1]
}