// https://leetcode.com/problems/search-a-2d-matrix/
package search_2d_matrix

func searchMatrix(matrix [][]int, target int) bool {
	// find possible row position
	var possiblePosition int = -1
	for k, row := range matrix {
		if row[0] <= target && k > possiblePosition {
			possiblePosition = k
		}
	}

	// iterate each number inside possible row position
	if possiblePosition > -1 {
		for _, n := range matrix[possiblePosition] {
			if n == target {
				return true
			}
		}
	}

	return false
}
