type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	encoded := ""
	for _, str := range strs {
		encoded = encoded + str + "€"
	}
	return encoded
}

func (s *Solution) Decode(encoded string) []string {
	out := make([]string, 0)
	curr := make([]rune, 0)
	for _, char := range encoded {
		if char == '€' {
			result := string(curr)
			out = append(out, result)
			curr = make([]rune, 0)
			continue
		}
		curr = append(curr, char)
	}
	return out
}
