package search_in_a_binary_search_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	for root != nil {
		if root.Val == val {
			return root
		}

		if root.Left == nil && root.Right == nil {
			return nil
		}

		if root.Val > val && root.Left != nil {
			root = root.Left
			continue
		}

		if root.Val < val && root.Right != nil {
			root = root.Right
			continue
		}

		return nil
	}

	return nil
}
