func permute(nums []int) [][]int {
    out := [][]int{}
    var parse func([]bool, []int)
    parse = func(used []bool, curr []int) {
        if len(curr) == len(nums) {
            out = append(out, curr)
            return
        }
        for idx, el := range used {
            if !el {
                used[idx] = true
                newCurr := append([]int{}, curr...)
                parse(used, append(newCurr, nums[idx]))
                // backtracking
                used[idx] = false
            }
        }
    }
    parse(make([]bool, len(nums)), []int{})
    return out
}