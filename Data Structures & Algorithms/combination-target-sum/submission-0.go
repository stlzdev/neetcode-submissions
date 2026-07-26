func combinationSum(nums []int, target int) [][]int {
    out := [][]int{}
    group := []int{}
    var csum func(int, int)
    csum = func(curr int, idx int) {
        if curr == target {
            out = append(out, append([]int{}, group...))
            return
        }
        if curr > target || idx >= len(nums) {
            return
        }
        group = append(group, nums[idx])
        csum(curr + nums[idx], idx)
        group = group[:len(group) - 1]
        csum(curr, idx + 1)
    } 
    csum(0, 0)
    return out
}
