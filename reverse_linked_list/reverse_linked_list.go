package reverse_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	newLn := new(ListNode)
	h := head
	i := 0

	for h != nil {
		hh := new(ListNode)
		*hh = *h
		hh.Next = nil

		if i == 0 {
			newLn = hh
		} else {
			hh.Next = newLn
			newLn = hh
		}

		h = h.Next
		i++
	}

	return newLn
}
