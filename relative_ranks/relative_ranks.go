package relative_ranks

import (
	"strconv"
)

func findRelativeRanks(score []int) []string {
	// make duplicate score with sorted position
	sortedScore := make([]int, len(score))
	copy(sortedScore, score)

	for i := 0; i < len(sortedScore)-1; i++ {
		for j := i + 1; j < len(sortedScore); j++ {
			if sortedScore[i] < sortedScore[j] {
				sortedScore[i], sortedScore[j] = sortedScore[j], sortedScore[i]
			}
		}
	}

	// map the result based on sorted score position
	var result []string
	for _, n := range score {
		for k, sN := range sortedScore {
			if n == sN {
				var rank string

				switch k {
				case 0:
					rank = "Gold Medal"
				case 1:
					rank = "Silver Medal"
				case 2:
					rank = "Bronze Medal"
				default:
					rank = strconv.Itoa(k + 1)
				}

				result = append(result, rank)
			}
		}
	}

	return result
}
