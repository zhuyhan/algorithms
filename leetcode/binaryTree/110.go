package binaryTree

import "math"

//110. 平衡二叉树
//https://leetcode-cn.com/problems/balanced-binary-tree/

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return math.Abs(height(root.Left)-height(root.Right)) < 2 && isBalanced(root.Left) && isBalanced(root.Right)

}

func height(root *TreeNode) float64 {
	if root == nil {
		return -1
	}
	l := height(root.Left)
	r := height(root.Right)
	return max(l, r) + 1
}

func max(l, r float64) float64 {
	if l > r {
		return l
	}
	return r
}
