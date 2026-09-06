func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	tot := len(nums1) + len(nums2)
	half := tot / 2
	a, b := nums1, nums2
	if len(a) > len(b) {
		a, b = b, a
	}
	l, r := 0, len(a)
	for {
		i := (l+r)/2
		j := half - i
		aLeft := math.MinInt
		if i > 0 {
			aLeft = a[i-1]
		}
		aRight := math.MaxInt
		if i < len(a) {
			aRight = a[i]
		}
		bLeft := math.MinInt
		if j > 0 {
			bLeft = b[j-1]
		}
		bRight := math.MaxInt
		if j < len(b) {
			bRight = b[j]
		}
		if aLeft <= bRight && bLeft <= aRight {
			if tot % 2 == 0 {
				return float64(max(aLeft, bLeft) + min(aRight, bRight)) / 2 
			}
			return float64(min(aRight, bRight))
		} else if aLeft > bRight {
			r = i - 1
		} else {
			l = i + 1
		}
	}
}
