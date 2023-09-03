// https://leetcode.com/problems/insertion-sort-list/

package insertion_sort_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	numbers := convertListNodeToSlice(head)
	for i := 0; i < len(numbers)-1; i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i] > numbers[j] {
				numbers[i], numbers[j] = numbers[j], numbers[i]
			}
		}
	}

	return convertSliceToListNode(numbers)
}

func convertListNodeToSlice(head *ListNode) []int {
	var numbers []int = []int{head.Val}
	node := head

	for node.Next != nil {
		numbers = append(numbers, node.Next.Val)
		node = node.Next
	}
	return numbers
}

func reverseSlice(numbers []int) []int {
	var newNumbers = []int{}
	for i := len(numbers) - 1; i >= 0; i-- {
		newNumbers = append(newNumbers, numbers[i])
	}
	return newNumbers
}

func convertSliceToListNode(numbers []int) *ListNode {
	var node *ListNode = nil
	for _, n := range reverseSlice(numbers) {
		node = &ListNode{
			Val:  n,
			Next: node,
		}
	}

	return node
}
