package insert_greatest_common_divisors_in_linked_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindGreatestCommonDivisors(t *testing.T) {
	testCases := []struct {
		name     string
		x        int
		y        int
		expected int
	}{
		{
			name:     "Test Find Greatest Common Divisors #1",
			x:        9,
			y:        6,
			expected: 3,
		},
		{
			name:     "Test Find Greatest Common Divisors #2",
			x:        81,
			y:        64,
			expected: 1,
		},
		{
			name:     "Test Find Greatest Common Divisors #3",
			x:        72,
			y:        64,
			expected: 8,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, findGreatestCommonDivisors(tc.x, tc.y))
		})
	}
}

func TestInsertGreatestCommonDivisors(t *testing.T) {
	testCases := []struct {
		name     string
		head     *ListNode
		expected *ListNode
	}{
		{
			name: "Test Insert Greatest Common Divisors #1",
			head: &ListNode{
				Val: 18,
				Next: &ListNode{
					Val: 6,
					Next: &ListNode{
						Val: 10,
						Next: &ListNode{
							Val: 3,
						},
					},
				},
			},
			expected: &ListNode{
				Val: 18,
				Next: &ListNode{
					Val: 6,
					Next: &ListNode{
						Val: 6,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val: 10,
								Next: &ListNode{
									Val: 1,
									Next: &ListNode{
										Val: 3,
									},
								},
							},
						},
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, insertGreatestCommonDivisors(tc.head))
			// cNode := insertGreatestCommonDivisors(tc.head)
			// for cNode != nil {
			// 	fmt.Println(">>", cNode.Val)
			// 	cNode = cNode.Next
			// }
		})
	}
}
