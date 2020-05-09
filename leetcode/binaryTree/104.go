package binaryTree

//104 二叉树的最大深度
//https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/

//法一：深度优先遍历
func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max104(maxDepth(root.Left), maxDepth(root.Right))
}
func max104(l, r int) int {
	if l > r {
		return l
	}
	return r
}

//法二：层序遍历；将每层的节点放在队列里
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var depth int
	queue := []*TreeNode{root}
	for ; len(queue) > 0; depth++ {
		var tmp []*TreeNode
		for _, v := range queue {
			if v.Left != nil {
				tmp = append(tmp, v.Left)
			}
			if v.Right != nil {
				tmp = append(tmp, v.Right)
			}
			queue = tmp
		}
	}
	return depth
}
