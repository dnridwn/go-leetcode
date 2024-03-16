package top_k_frequent_elements

func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	maxFreq := 0
	for _, num := range nums {
		count[num]++
		if count[num] > maxFreq {
			maxFreq = count[num]
		}
	}

	buckets := make([][]int, maxFreq+1)
	for num, freq := range count {
		buckets[freq] = append(buckets[freq], num)
	}

	result := make([]int, 0, k)
	for freq := maxFreq; freq > 0 && len(result) < k; freq-- {
		result = append(result, buckets[freq]...)
	}

	return result[:k]
}
