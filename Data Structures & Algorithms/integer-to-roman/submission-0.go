func intToRoman(num int) string {
    vals := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	syms := []string{"M", "CM", "D", "CD", "C", "XC", 
	"L", "XL", "X", "IX", "V", "IV", "I"}
	out := ""
	for i := 0; i < len(vals); i++ {
		for num >= vals[i] {
			out += syms[i]
			num -= vals[i]
		}
	}
	return out
}
