// https://leetcode.com/problems/add-two-numbers
package add_two_numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		lastL1, lastL2 *ListNode = l1, l2
		tempN          int       = 0
		result         *ListNode = nil
		lastResult     *ListNode = nil
	)

	for lastL1 != nil || lastL2 != nil || tempN > 0 {
		l1N := 0
		if lastL1 != nil {
			l1N = lastL1.Val
		}

		l2N := 0
		if lastL2 != nil {
			l2N = lastL2.Val
		}

		lN := l1N + l2N

		if tempN > 0 {
			lN += tempN
			tempN = 0
		}

		if lN > 9 {
			tempN = lN / 10
			lN = lN % 10
		}

		rs := &ListNode{
			Val: lN,
		}

		if result == nil {
			result = rs
		} else {
			lastResult.Next = rs
		}

		if lastL1 != nil {
			lastL1 = lastL1.Next
		}

		if lastL2 != nil {
			lastL2 = lastL2.Next
		}

		lastResult = rs
	}

	return result
}
