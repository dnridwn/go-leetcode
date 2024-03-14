package remove_duplicates_from_sorted_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	if head.Val == head.Next.Val {
		head.Next = head.Next.Next
		return deleteDuplicates(head)
	}

	head.Next = deleteDuplicates(head.Next)
	return head
}
