func dailyTemperatures(temperatures []int) []int {
    stack := []int{}
    result := make([]int, len(temperatures))
    for idx, el := range temperatures {
        for len(stack) > 0 && temperatures[stack[len(stack) - 1]] < el {
            hidx := stack[len(stack) - 1]
            stack = stack[:len(stack) - 1]
            result[hidx] = idx - hidx
        }
        stack = append(stack, idx)
    }
    return result
}
