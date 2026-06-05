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
			hrs += (val + m - 1) / m
			if hrs > h {
				break
			}
		}
		if hrs > h {
			l = m + 1
		} else {
			r = m - 1
		} 
	}
	return l
}
