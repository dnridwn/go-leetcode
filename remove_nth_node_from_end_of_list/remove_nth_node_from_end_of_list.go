// https://leetcode.com/problems/remove-nth-node-from-end-of-list/

package remove_nth_node_from_end_of_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	nodes := make([]*ListNode, 0)
	node := head
	for node != nil {
		nodes = append(nodes, node)
		node = node.Next
	}

	if n <= len(nodes) {
		kTarget := len(nodes) - n
		if kTarget > -1 {
			if kTarget == 0 {
				head = nodes[kTarget].Next
			} else {
				nodes[kTarget-1].Next = nil
				if nodes[kTarget].Next != nil {
					nodes[kTarget-1].Next = nodes[kTarget].Next
				}
			}
		}
	}

	return head
}
