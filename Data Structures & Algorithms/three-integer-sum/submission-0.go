import "slices"

func threeSum(nums []int) [][]int {
	out := make([][]int, 0)
	seen := make(map[[3]int] bool)
	slices.Sort(nums)
	for idx, num := range nums {
		if num > 0 {
			break
		}
		comp := -num
		lpt := idx + 1
		rpt := len(nums) - 1
		for lpt < rpt {
			nlpt := nums[lpt]
			nrpt := nums[rpt]
			sum := nums[lpt] + nums[rpt]
			if sum == comp {
				key := [3]int{num, nlpt, nrpt}
				if !seen[key] {
					seen[key] = true
					out = append(out, []int{num, nlpt, nrpt})
				}
				lpt++
				rpt--
			} else if sum < comp {
				lpt++
			} else {
				rpt--
			}
		}
	}
	return out
}
