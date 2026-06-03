func evalRPN(tokens []string) int {
	stack := []int{}
	oper := map[string](func(int, int)int){
		"+": func(a, b int) int {return a + b},
		"-": func(a, b int) int {return a - b},
		"*": func(a, b int) int {return a * b},
		"/": func(a, b int) int {return a / b},
	}
	for _, str := range tokens {
		if fn, exists := oper[str]; exists {
			val1 := stack[len(stack) - 2]
			val2 := stack[len(stack) - 1]
			stack = stack[:len(stack) - 2]
			newval := fn(val1, val2)
			stack = append(stack, newval)
		} else {
			num, _ := strconv.Atoi(str)
			stack = append(stack, num)
		}
	}
	return stack[len(stack) - 1]
}
