package search_in_a_binary_search_tree

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchBST(t *testing.T) {
	tc := []struct {
		root     *TreeNode
		val      int
		expected *TreeNode
	}{
		{
			root: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
				Right: &TreeNode{
					Val: 7,
				},
			},
			val: 2,
			expected: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
		},
		{
			root: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
				Right: &TreeNode{
					Val: 7,
				},
			},
			val:      5,
			expected: nil,
		},
	}

	for k, tt := range tc {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			assert.Equal(t, tt.expected, searchBST(tt.root, tt.val))
		})
	}
}
