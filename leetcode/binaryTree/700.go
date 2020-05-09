package binaryTree

func searchBST(root *TreeNode, val int) *TreeNode {
	for root != nil {
		if root.Val == val {
			return root
		}
		if root.Val > val {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return nil
}
