func twoSum(nums []int, target int) []int {
	set := make(map[int] int)
    for i, val := range nums {
		comp := target - val
		if ci, exist := set[comp]; exist {
			return []int{ci, i}
		}
		set[val] = i 
	}
	return []int{}
}
