package kids_with_the_greetest_number_of_candies

import "slices"

func kidsWithCandies(candies []int, extraCandies int) []bool {
	m := slices.Max(candies)
	r := make([]bool, len(candies))

	for k, c := range candies {
		r[k] = c+extraCandies >= m
	}

	return r
}
