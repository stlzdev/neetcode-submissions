func productExceptSelf(nums []int) []int {
	out := make([]int, len(nums))
	for j := range out {
		out[j] = 1
	}
	// left pass
	running := 1
	for i, num := range nums {
		out[i] = running
		running = out[i] * num
	}
	// right pass
	running = 1
	for k := len(nums) - 1; k >= 0; k-- {
		out[k] *= running
		running *= nums[k]
	}
	return out
}