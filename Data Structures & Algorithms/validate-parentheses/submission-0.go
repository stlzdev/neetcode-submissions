func isValid(s string) bool {
	stack := []rune{}
	pair := map[rune]rune {
		')' : '(',
		'}' : '{',
		']': '[',
	}
	for _, r := range s {
		if r == '(' || r == '{' || r == '[' {
			stack = append(stack, r)
		} else if len(stack) > 0 && stack[len(stack) - 1] == pair[r] {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}