package reverse_linked_list

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseList(t *testing.T) {
	tc := []struct {
		head     *ListNode
		expected *ListNode
	}{
		{
			head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 5,
							},
						},
					},
				},
			},
			expected: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val: 1,
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
				},
			},
			expected: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
				},
			},
		},
	}

	for k, tt := range tc {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			assert.Equal(t, tt.expected, reverseList(tt.head))
		})
	}
}
