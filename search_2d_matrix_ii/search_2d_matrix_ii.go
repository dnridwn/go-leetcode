// https://leetcode.com/problems/search-a-2d-matrix-ii/
package search_2d_matrix_ii

func searchMatrix(matrix [][]int, target int) bool {
	var (
		possibleXs []int
		possibleYs []int
	)

	// find possible X's
	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] <= target {
			possibleXs = append([]int{i}, possibleXs...)
		}
	}

	// find possible Y's
	for i := 0; i < len(matrix[0]); i++ {
		if matrix[0][i] <= target {
			possibleYs = append([]int{i}, possibleYs...)
		}
	}

	// find the target
	for _, i := range possibleXs {
		for _, j := range possibleYs {
			if matrix[i][j] == target {
				return true
			}
		}
	}

	return false
}
