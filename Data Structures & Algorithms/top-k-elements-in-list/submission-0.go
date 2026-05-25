func topKFrequent(nums []int, k int) []int {
	freqs := make(map[int]int)
	for _, num := range nums {
		freqs[num]++
	}
	n := len(nums)
	buckets := make([][]int, n+1)
	for num, count := range freqs {
		buckets[count] = append(buckets[count], num)
	}
	result := make([]int, 0, k)
	for i := n; i >= 0 && len(result) < k; i-- {
		if len(buckets[i]) > 0 {
			result = append(result, buckets[i]...)
		}
	}
	return result[:k]
}
