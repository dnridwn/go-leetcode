package unique_number_of_occurrences

func uniqueOccurrences(arr []int) bool {
	freq := make(map[int]int)
	for _, n := range arr {
		if _, ok := freq[n]; !ok {
			freq[n] = 0
		}

		freq[n]++
	}

	freqExist := make(map[int]bool)
	for _, f := range freq {
		if _, ok := freqExist[f]; ok {
			return false
		}

		freqExist[f] = true
	}

	return true
}
