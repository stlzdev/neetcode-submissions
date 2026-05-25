type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	encoded := ""
	for _, str := range strs {
		encoded = encoded + strconv.Itoa(len(str)) + "*" + str
	}
	return encoded
}

func (s *Solution) Decode(encoded string) []string {
	out := make([]string, 0)
	num := make([]rune, 0)
	curr := make([]rune, 0)
	numlen := 0
	for _, char := range encoded {
		if numlen > 0 {
			curr = append(curr, char)
			numlen--
			if len(curr) > 0 && numlen == 0 {
				out = append(out, string(curr))
				curr = make([]rune, 0)
			}
		} else if char >= '0' && char <= '9' {
			num = append(num, char)
		} else if char == '*' {
			numlen2, err := strconv.Atoi(string(num))
			numlen = numlen2
			if err != nil {
				print("invalid input")
			}
			num = make([]rune, 0)
			if numlen == 0 {
				out = append(out, "")
			}
		}
	}
	return out
}
