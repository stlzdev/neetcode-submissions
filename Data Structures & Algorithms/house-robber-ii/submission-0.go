func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	} else if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	dp1 := make([]int, len(nums)-1)
	dp2 := make([]int, len(nums)-1)
	dp1[0] = nums[0]; dp1[1] = max(nums[0], nums[1])
	dp2[0] = nums[1]; dp2[1] = max(nums[1], nums[2])
	for i := 2; i < len(nums)-1; i++ {
		dp1[i] = max(dp1[i-2] + nums[i], dp1[i-1])
		dp2[i] = max(dp2[i-2] + nums[i+1], dp2[i-1])
	}
	return max(dp1[len(nums)-2], dp2[len(nums)-2])
}
