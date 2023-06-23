package sort_people

func sortPeople(names []string, heights []int) []string {
	// mapping the [height]people
	mapping := make(map[int]string)
	for k, v := range names {
		mapping[heights[k]] = v
	}

	// sorting the height
	for i := 0; i < len(heights)-1; i++ {
		for j := i + 1; j < len(heights); j++ {
			if heights[j] > heights[i] {
				heights[i], heights[j] = heights[j], heights[i]
			}
		}
	}

	// append the result
	finalResult := make([]string, 0)
	for _, v := range heights {
		finalResult = append(finalResult, mapping[v])
	}

	return finalResult
}
