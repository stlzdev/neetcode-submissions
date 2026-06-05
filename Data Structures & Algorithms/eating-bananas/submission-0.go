func minEatingSpeed(piles []int, h int) int {
	max := piles[0]
	for _, num := range piles {
		if num > max {
			max = num
		}
	}
	l, r := 1, max
	for l <= r {
		m := l + (r-l) / 2
		hrs := 0
		for _, val := range piles {
			hrs += int(math.Ceil(float64(val) / float64(m)))
		}
		if hrs > h {
			l = m + 1
		} else {
			r = m - 1
		} 
	}
	return l
}
