package rotate

import "fmt"

func rotate(matrix [][]int) [][]int {
	history := map[string]bool{}

	for i, x := range matrix {
		for j, y := range x {
			if _, ok := history[fmt.Sprintf("%d*%d", i, j)]; ok {
				continue
			}

			jj := len(matrix) - 1 - i
			yy := matrix[j][jj]

			matrix[i][j] = yy
			matrix[j][jj] = y

			history[fmt.Sprintf("%d*%d", i, j)] = true
			history[fmt.Sprintf("%d*%d", j, jj)] = true
		}
	}

	return matrix
}
