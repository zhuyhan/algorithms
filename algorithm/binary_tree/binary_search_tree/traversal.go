package binary_search_tree

import (
	"container/list"
	"fmt"
)

//前序遍历-递归
func preOrderRecursion(root *Node) {
	if root == nil {
		return
	}
	fmt.Println(root.element)
	preOrderRecursion(root.left)
	preOrderRecursion(root.right)
}

//前序遍历-非递归
func preOrderNoRecursion(root *Node) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	stack := list.New()

	for root != nil || stack.Len() != 0 {
		for root != nil {
			result = append(result, root.element)
			stack.PushBack(root)
			root = root.left

		}
		if stack.Len() != 0 {
			v := stack.Back()
			root = v.Value.(*Node)
			result = append(result, root.element)
			root = root.right
			stack.Remove(v)
		}
	}
	return result
}

//中序遍历-递归
func inOrderRecursion(root *Node) {
	if root == nil {
		return
	}
	inOrderRecursion(root.left)
	fmt.Println(root.element)
	inOrderRecursion(root.right)
}

//后序遍历-递归
func postOrderRecursion(root *Node) {
	if root == nil {
		return
	}
	postOrderRecursion(root.left)
	postOrderRecursion(root.right)
	fmt.Println(root.element)
}
