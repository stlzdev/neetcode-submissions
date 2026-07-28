import (
	"slices"
)
func subsetsWithDup(nums []int) [][]int {
	out := [][]int{}
	slices.Sort(nums)
	var add func([]int, int)
	add = func(curr []int, idx int) {
		if idx == len(nums) {
			newCurr := append([]int{}, curr...)
			out = append(out, newCurr)
			return
		} 
		// skip current index element
		nxt := idx + 1
		for nxt < len(nums) {
			if nums[nxt] != nums[idx] {
				break
			}
			nxt += 1
		}
		add(curr, nxt)
		// include current index element
		curr = append(curr, nums[idx])
		add(curr, idx + 1)
		curr = curr[:len(curr) - 1]
	}
	add([]int{}, 0)
	return out
}
