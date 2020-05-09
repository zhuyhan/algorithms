package binary_search_tree

import (
	"fmt"
)

type Node struct {
	element             int
	left, right, parent *Node
}

type Tree struct {
	root *Node
	size int
}

func NewTree() *Tree {
	return &Tree{}
}

func (t *Tree) Add(element int) bool {
	//element如果是interface{}类型；要判断nil
	//根结点判断
	if t.root == nil {
		t.root = &Node{element: element}
		t.size++
		return true
	}
	//其他结点
	return t.add(element)
}

//add 方法一：非递归添加
func (t *Tree) add(element int) bool {
	node := t.root
	var parent *Node
	var compare int
	//找到父节点和插入位置
	for node != nil {
		parent = node
		if node.element > element {
			node = node.left
			compare = 1
		} else if node.element < element {
			node = node.right
			compare = -1
		} else {
			//相等直接返回
			return true
		}
	}
	//插入父节点的哪个位置
	if compare == 1 {
		parent.left = &Node{
			element: element,
			parent:  parent,
		}
	} else if compare == -1 {
		parent.right = &Node{
			element: element,
			parent:  parent,
		}
	}
	t.size++
	return true
}

//add2 方法二：递归方法
func (t *Tree) add2(element int) bool {
	return true
}

//PreOrderTraversal 前序遍历
func (t *Tree) PreOrderTraversal() bool {
	preOrderRecursion(t.root)
	return false
}

//InOrderTraversal 中序遍历
func (t *Tree) InOrderTraversal() bool {
	inOrderRecursion(t.root)
	return false
}

//PostOrderTraversal 后序遍历
func (t *Tree) PostOrderTraversal() bool {
	postOrderRecursion(t.root)
	return false
}

//LevelOrderTraversal 层序遍历
func (t *Tree) LevelOrderTraversal() []int {
	var list []int
	if t.root == nil {
		return list
	}
	queue := make([]*Node, 0)
	queue = append(queue, t.root)

	for len(queue) != 0 {
		tmp := make([]*Node, 0)
		for _, v := range queue {
			list = append(list, v.element)
			if v.left != nil {
				tmp = append(tmp, v.left)
			}
			if v.right != nil {
				tmp = append(tmp, v.right)
			}
		}
		queue = tmp
	}
	return list
}

//Height 树的高度（层序遍历）
func (t *Tree) Height() int {
	if t.root == nil {
		return 0
	}
	var height int
	queue := make([]*Node, 0)
	queue = append(queue, t.root)

	for len(queue) != 0 {
		tmp := make([]*Node, 0)
		height++
		for _, v := range queue {
			if v.left != nil {
				tmp = append(tmp, v.left)
			}
			if v.right != nil {
				tmp = append(tmp, v.right)
			}
		}
		queue = tmp
	}
	return height
}

//Height 树的高度（递归）
func (t *Tree) Height2(root *Node) int {
	if root == nil {
		return 0
	}
	var height int
	h1 := t.Height2(root.left)
	h2 := t.Height2(root.right)
	if h1 > h2 {
		height = h1
	} else {
		height = h2
	}
	return height + 1
}

//IsComplete 判断树是否为完全二叉树
func (t *Tree) IsComplete() bool {
	if t.root == nil {
		return false
	}
	queue := make([]*Node, 0)
	queue = append(queue, t.root)
	var leaf bool
	for len(queue) != 0 {
		tmp := make([]*Node, 0)

		for _, v := range queue {
			if leaf && isLeaf(v) != true {
				return false
			}
			if hasTwoChildren(v) == true {
				tmp = append(tmp, v.left)
				tmp = append(tmp, v.right)
			} else if v.left == nil && v.right != nil {
				return false
			} else {
				leaf = true
				if v.left != nil {
					tmp = append(tmp, v.left)
				}
			}

		}
		queue = tmp
	}
	return true
}

//isLeaf 判断是否为叶子节点
func isLeaf(leaf *Node) bool {
	if leaf.right == nil && leaf.left == nil {
		return true
	}
	return false
}

//hasTwoChildren 判断是否有左右子树
func hasTwoChildren(root *Node) bool {
	if root.right != nil && root.left != nil {
		return true
	}
	return false
}

//判断前驱节点
func (t *Tree) Predecessor(node *Node) *Node {
	fmt.Println(node)
	if node == nil {
		return nil
	}

	var pre *Node = nil
	if node.left != nil {
		tmp := node.left
		for tmp != nil {
			pre = tmp
			tmp = tmp.right
		}
	}
	if node.left == nil && node.parent != nil {
		parent := node.parent
		for parent != nil {
			if parent.element < node.element {
				pre = parent
			}
			parent = parent.parent
		}
	}
	fmt.Println(pre)
	return pre
}

//判断后驱节点
func (t *Tree) Successor(node *Node) *Node {
	fmt.Println(node)
	if node == nil {
		return nil
	}

	var pre *Node = nil
	if node.right != nil {
		tmp := node.right
		for tmp != nil {
			pre = tmp
			tmp = tmp.left
		}
	}
	if node.right == nil && node.parent != nil {
		parent := node.parent
		for parent != nil {
			if parent.element > node.element {
				pre = parent
				return pre
			}
			parent = parent.parent
		}
	}
	fmt.Println(pre)
	return pre
}

//查找结点
func (t *Tree) GetNode(element int) *Node {
	node := t.root
	for node != nil {
		if node.element == element {
			return node
		}
		if node.element > element {
			node = node.left
		}
		if node.element < element {
			node = node.right
		}
	}
	return nil
}

//删除结点
func (t *Tree) Remove(element int) bool {
	//获取节点
	node := t.GetNode(element)
	if node == nil {
		return false
	}
	//判断节点的度为2
	if hasTwoChildren(node) == true {
		//找到后继节点
		s := t.Successor(node)
		fmt.Println("找到后继节点：", s)
		//用后继节点的值替换
		node.element = s.element
		//删除后继节点
		node = s
	}
	//判断节点的度为1或0
	var replacement *Node
	if node.left != nil {
		replacement = node.left
	} else {
		replacement = node.right
	}
	if replacement != nil {
		//度为1
		replacement.parent = node.parent
		if node.parent == nil {
			//node为根节点
			t.root = replacement
		} else if node == node.parent.left {
			node.parent.left = replacement
		} else {
			node.parent.right = replacement
		}
	} else if node.parent == nil {
		//node是根节点
		t.root = nil
	} else {
		//node是叶子节点
		if node.parent.left == node {
			node.parent.left = nil
		} else {
			node.parent.right = nil
		}
	}
	return true
}
