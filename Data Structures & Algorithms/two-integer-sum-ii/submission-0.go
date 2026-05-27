func twoSum(numbers []int, target int) []int {
	hmap := make(map[int]int)
	for idx, num := range numbers {
		comp := target - num
		cidx, exist := hmap[comp]
		if exist {
			return []int{cidx + 1, idx + 1}
		}
		hmap[num] = idx 
	}
	return []int{}
}