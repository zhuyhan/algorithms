package binaryTree

//94. 二叉树的中序遍历
//https://leetcode-cn.com/problems/binary-tree-inorder-traversal/

func inorderTraversal(root *TreeNode) []int {
	list := make([]int, 0)
	list = inorder(root, list)
	return list
}

func inorder(root *TreeNode, list []int) []int {
	if root == nil {
		return list
	}
	list = inorder(root.Left, list)
	list = append(list, root.Val)
	list = inorder(root.Right, list)
	return list
}
