package insert_greatest_common_divisors_in_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func findGreatestCommonDivisors(x int, y int) int {
	var smN int
	if x <= y {
		smN = x
	} else {
		smN = y
	}

	var gcd int
	for i := 1; i <= smN; i++ {
		if x%i == 0 && y%i == 0 {
			gcd = i
		}
	}

	return gcd
}

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	cNode := head
	for cNode.Next != nil {
		x := cNode.Val
		y := cNode.Next.Val

		nNode := &ListNode{
			Val:  findGreatestCommonDivisors(x, y),
			Next: cNode.Next,
		}

		cNode.Next = nNode
		cNode = cNode.Next.Next
	}

	return head
}
