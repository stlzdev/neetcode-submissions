import (
	"slices"
)

func combinationSum2(candidates []int, target int) [][]int {
    out := [][]int{}
    group := []int{} 
    slices.Sort(candidates)
    var csum func(int, int)
    csum = func(curr int, idx int) {
        if curr == target {
            out = append(out, append([]int{}, group...))
            return
        }
        if curr > target || idx >= len(candidates) {
            return
        }
        group = append(group, candidates[idx])
        csum(curr + candidates[idx], idx + 1)
        group = group[:len(group) - 1]
        // find next different index
        el := candidates[idx]
        nxt := idx + 1
        for nxt < len(candidates) && candidates[nxt] == el {
            nxt++
        } 
        if nxt < len(candidates) {
            csum(curr, nxt)
        }
    }
    csum(0, 0)
    return out
}
