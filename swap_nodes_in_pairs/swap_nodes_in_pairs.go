// https://leetcode.com/problems/swap-nodes-in-pairs/

package swap_nodes_in_pairs

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var (
		currNode *ListNode = head
		newHead  *ListNode = head.Next
	)

	for currNode != nil && currNode.Next != nil {
		temp := currNode.Next
		currNode.Next = currNode.Next.Next
		temp.Next = currNode
		currNode = currNode.Next

		if currNode != nil && currNode.Next != nil {
			temp.Next.Next = currNode.Next
		}
	}

	return newHead
}
