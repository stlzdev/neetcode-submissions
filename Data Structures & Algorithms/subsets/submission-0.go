func subsets(nums []int) [][]int {
    var out [][]int
    var backtrack func(int, []int)
    backtrack = func(i int, path []int) {
        if i == len(nums) {
            out = append(out, append([]int{}, path...))
            return
        }
        newPath1 := append([]int{}, path...)
        backtrack(i+1, newPath1)
        newPath2 := append([]int{}, path...)
        newPath2 = append(newPath2, nums[i])
        backtrack(i+1, newPath2)
    }
    backtrack(0, []int{})
    return out
}
