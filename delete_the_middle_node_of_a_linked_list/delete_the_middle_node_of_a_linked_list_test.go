package delete_the_middle_node_of_a_linked_list

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteMiddle(t *testing.T) {
	testCases := []struct {
		head     *ListNode
		expected *ListNode
	}{
		{
			head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 7,
							Next: &ListNode{
								Val: 1,
								Next: &ListNode{
									Val: 2,
									Next: &ListNode{
										Val: 6,
									},
								},
							},
						},
					},
				},
			},
			expected: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 1,
							Next: &ListNode{
								Val: 2,
								Next: &ListNode{
									Val: 6,
								},
							},
						},
					},
				},
			},
		},
		{
			head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
						},
					},
				},
			},
			expected: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 4,
					},
				},
			},
		},
		{
			head: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
				},
			},
			expected: &ListNode{
				Val: 2,
			},
		},
	}

	for k, tc := range testCases {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			assert.Equal(t, tc.expected, deleteMiddle(tc.head))
		})
	}
}
