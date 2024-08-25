package odd_even_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var (
		odds, lastOdd, evens, lastEven *ListNode
		isOdd                          bool = false
	)

	for head != nil {
		cHead := *head
		cHead.Next = nil

		if isOdd {
			if odds == nil {
				odds = &cHead
				lastOdd = odds
			} else {
				lastOdd.Next = &cHead
				lastOdd = lastOdd.Next
			}
		} else {
			if evens == nil {
				evens = &cHead
				lastEven = evens
			} else {
				lastEven.Next = &cHead
				lastEven = lastEven.Next
			}
		}

		head = head.Next
		isOdd = !isOdd
	}

	lastEven.Next = odds
	return evens
}
