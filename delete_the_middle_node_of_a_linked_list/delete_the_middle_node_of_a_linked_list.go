package delete_the_middle_node_of_a_linked_list

import "math"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteMiddle(head *ListNode) *ListNode {
	l := []*ListNode{}
	r := head

	for r != nil {
		l = append(l, r)
		r = r.Next
	}

	if len(l) == 1 {
		return nil
	}

	m := int(math.Floor(float64(len(l)) / float64(2)))

	if m == len(l)-1 {
		l[m-1].Next = nil
	} else {
		l[m-1].Next = l[m+1]
	}

	return head
}
