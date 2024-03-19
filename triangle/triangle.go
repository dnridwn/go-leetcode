package triangle

import (
	"fmt"
)

func minimumTotal(triangle [][]int) int {
	memo := map[string]int{}
	i := dive(triangle, 0, 0, memo)
	return i
}

func dive(triangle [][]int, row int, col int, memo map[string]int) int {
	key := fmt.Sprintf("%d,%d", row, col)
	if val, ok := memo[key]; ok {
		return val
	}

	c := triangle[row][col]

	if row == len(triangle)-1 {
		return c
	}

	left := dive(triangle, row+1, col+0, memo)
	right := dive(triangle, row+1, col+1, memo)

	min := left
	if right < left {
		min = right
	}
	min += c

	memo[key] = min
	return min
}
