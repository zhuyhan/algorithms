package binaryTree

//617. 合并二叉树
//https://leetcode-cn.com/problems/merge-two-binary-trees/

//法一：递归
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	t1.Val += t2.Val
	t1.Left = mergeTrees(t1.Left, t2.Left)
	t1.Right = mergeTrees(t1.Right, t2.Right)
	return t1
}

//法二：中序遍历

func reverseLeftWords(s string, n int) string {
	if n < 1 || n >= len(s) {
		return s
	}
	return s[n:] + s[:n]
}
