package same_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSameTree(t *testing.T) {
	testCases := []struct {
		name     string
		pTree    *TreeNode
		qTree    *TreeNode
		expected bool
	}{
		{
			name: "Test Identical",
			pTree: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			qTree: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			expected: true,
		},
		{
			name: "Test Not Identical (1)",
			pTree: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
			},
			qTree: &TreeNode{
				Val:  1,
				Left: nil,
				Right: &TreeNode{
					Val: 2,
				},
			},
			expected: false,
		},
		{
			name: "Test Not Identical (2)",
			pTree: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 1,
				},
			},
			qTree: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
			expected: false,
		},
		{
			name:     "Test Both Nil",
			pTree:    nil,
			qTree:    nil,
			expected: true,
		},
		{
			name:  "Test One of Them is Nil",
			pTree: nil,
			qTree: &TreeNode{
				Val: 0,
			},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, isSameTree(tc.pTree, tc.qTree))
		})
	}
}
