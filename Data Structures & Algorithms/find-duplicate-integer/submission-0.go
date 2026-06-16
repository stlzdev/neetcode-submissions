func findDuplicate(nums []int) int {
	for i := 1; i <= len(nums); i++ {
		repeat := false
		for _ , num := range nums {
			if num == i && repeat {
				return num
			} else if num == i {
				repeat = true
			}
		}
	}
	return 0
}
