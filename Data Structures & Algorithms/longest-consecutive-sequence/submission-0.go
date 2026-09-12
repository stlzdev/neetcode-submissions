func longestConsecutive(nums []int) int {
	set := make(map[int]struct{}, len(nums))
	for _, n := range nums {
		set[n] = struct{}{}
	}
	maxl := 0
	for v, _ := range set {
		if _, exist := set[v-1]; exist {
			continue
		}
		curr := 1
		i := v
		for {
			if _, exist := set[i+1]; !exist {
				break
			}
			curr++
			i++
		}
		maxl = max(maxl, curr)
	}
	return maxl
}
