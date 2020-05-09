package AVL

import (
	"errors"
	"fmt"
)

type Node struct {
	element, height     int64
	left, right, parent *Node
}

type Tree struct {
	root *Node
	size int
}

func NewTree() *Tree {
	return &Tree{}
}

//Add 添加
func (t *Tree) Add(element int64) bool {
	//element如果是interface{}类型；要判断nil
	//根结点判断
	if t.root == nil {
		t.root = &Node{element: element, height: 1}
		t.size++
		return true
	}
	//添加节点
	node := t.add(element)
	parent := node.parent
	for parent != nil {
		//调整树结构
		if parent.isBalance() {
			//更新高度（父节点，祖宗节点）
			parent.updateHeight()
		} else {
			//恢复平衡
			t.reBalance(parent)
			break
		}
		parent = parent.parent
	}

	return true
}

//Remove 删除节点
func (t *Tree) Remove(element int64) bool {
	node, err := t.remove(element)
	fmt.Println(node)
	if err == nil {
		return true
	}
	return false
}

//LevelOrderTraversal 层序遍历
func (t *Tree) LevelOrderTraversal() []int64 {
	var list []int64
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

func (t *Tree) add(element int64) *Node {
	var node *Node //要插入的节点
	root := t.root
	var parent *Node
	var compare int
	//找到父节点和插入位置
	for root != nil {
		parent = root
		if root.element > element {
			root = root.left
			compare = 1
		} else if root.element < element {
			root = root.right
			compare = -1
		} else {
			//相等直接返回
			return root
		}
	}
	node = &Node{
		element: element,
		parent:  parent,
		height:  1,
	}
	//插入父节点的哪个位置
	if compare == 1 {
		parent.left = node

	} else if compare == -1 {
		parent.right = node
	}
	t.size++
	return node
}

func (n *Node) isBalance() bool {
	factor := n.balanceFactor()
	if factor > 1 || factor < -1 {
		return false
	}
	return true
}

//balanceFactor 判断一个节点的平衡因子
func (n *Node) balanceFactor() int64 {
	//左子树高度
	var left, right int64
	if n.left != nil {
		left = n.left.height
	}
	if n.right != nil {
		right = n.right.height
	}
	return left - right
}

//updateHeight 更新高度
func (n *Node) updateHeight() {
	var l, r int64
	if n.left != nil {
		l = n.left.height
	}
	if n.right != nil {
		r = n.right.height
	}
	if l < r {
		n.height = 1 + r
	} else {
		n.height = 1 + l
	}
	return
}

//reBalance 恢复平衡
func (t *Tree) reBalance(n *Node) {
	parent := n.tallerChild()
	child := parent.tallerChild()
	if parent.isLeftChild() {
		if child.isLeftChild() { //LL
			t.rotateRight(n)
		} else { //LR
			t.rotateLeft(parent)
			t.rotateRight(n)
		}
	} else {
		if child.isLeftChild() { //RL
			t.rotateRight(parent)
			t.rotateLeft(n)
		} else { //RR
			t.rotateLeft(n)
		}
	}
}

//rotateRight 右旋转
func (t *Tree) rotateRight(grand *Node) {
	parent := grand.left
	child := parent.right

	grand.left = child
	parent.right = grand

	t.afterRotate(grand, parent, child)
}

//rotateLeft 左旋
func (t *Tree) rotateLeft(grand *Node) {
	parent := grand.right
	child := parent.left

	grand.right = child
	parent.left = grand

	t.afterRotate(grand, parent, child)
}

func (t *Tree) afterRotate(grand, parent, child *Node) {
	//让parent为子树的根节点
	parent.parent = grand.parent
	if grand.isLeftChild() {
		grand.parent.left = parent
	} else if grand.isRightChild() {
		grand.parent.right = parent
	} else {
		//根节点
		t.root = parent
		//parent.parent = nil
	}
	//更新child的parent
	if child != nil {
		child.parent = grand
	}
	grand.parent = parent
	//更新节点高度(从低到高)
	grand.updateHeight()
	parent.updateHeight()
}

//tallerChild 获取高度更高的子节点
func (n *Node) tallerChild() *Node {
	var l, r int64
	if n.left != nil {
		l = n.left.height
	}
	if n.right != nil {
		r = n.right.height
	}
	if l > r {
		return n.left
	}
	return n.right
}

//判断自己是左子树
func (n *Node) isLeftChild() bool {
	if n != nil && n.parent != nil && n.parent.left == n {
		return true
	}
	return false
}

//判断自己是左子树
func (n *Node) isRightChild() bool {
	if n != nil && n.parent != nil && n.parent.right == n {
		return true
	}
	return false
}

func (t *Tree) remove(element int64) (*Node, error) {
	//获取节点
	node := t.GetNode(element)
	if node == nil {
		return nil, errors.New("未找到节点")
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
	//删除节点后处理
	t.afterRemove(node)
	return node, nil
}

func (t *Tree) afterRemove(n *Node) {
	parent := n.parent
	for parent != nil {
		//调整树结构
		if parent.isBalance() {
			//更新高度（父节点，祖宗节点）
			parent.updateHeight()
		} else {
			//恢复平衡
			t.reBalance(parent)
		}
		parent = parent.parent
	}
}

//GetNode 查找结点
func (t *Tree) GetNode(element int64) *Node {
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

//hasTwoChildren 判断是否有左右子树
func hasTwoChildren(root *Node) bool {
	if root.right != nil && root.left != nil {
		return true
	}
	return false
}

//Predecessor 判断前驱节点
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
	return pre
}

//Successor 判断后驱节点
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
