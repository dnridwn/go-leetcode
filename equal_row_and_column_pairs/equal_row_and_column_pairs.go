package equal_row_and_column_pairs

import (
	"strconv"
	"strings"
)

func equalPairs(grid [][]int) int {
	m := map[string]int{}
	c := 0

	for i := 0; i < len(grid); i++ {
		m[strings.Join(convertArrIntToStr(grid[i]), "|")]++
	}

	for i := 0; i < len(grid); i++ {
		r := []int{}
		for j := 0; j < len(grid); j++ {
			r = append(r, grid[j][i])
		}

		if f, ok := m[strings.Join(convertArrIntToStr(r), "|")]; ok {
			c += f
		}
	}

	return c
}

func convertArrIntToStr(nums []int) []string {
	r := []string{}
	for _, n := range nums {
		r = append(r, strconv.Itoa(n))
	}
	return r
}
