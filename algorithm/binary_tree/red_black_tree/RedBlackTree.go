package red_black_tree

import (
	"errors"
	"fmt"
)

const (
	RED   = false
	BLACK = true
)

type Node struct {
	color               bool
	element             int64
	left, right, parent *Node
}

type NodeValue struct {
	element int64
	color   bool
}

type Tree struct {
	root *Node
	size int
}

func NewTree() *Tree {
	return &Tree{}
}

//color 染色
func color(n *Node, c bool) *Node {
	if n == nil {
		return n
	}
	n.color = c
	return n
}

//red 染成红色
func red(n *Node) *Node {
	return color(n, RED)
}

//black染成黑色
func black(n *Node) *Node {
	return color(n, BLACK)
}

//colorOf 返回节点的颜色
func (n *Node) colorOf() bool {
	if n == nil {
		return BLACK
	}
	return n.color
}

//isBlack 节点是否为黑色
func (n *Node) isBlack() bool {
	return n.colorOf() == BLACK
}

//isRed 节点是否为红色
func (n *Node) isRed() bool {
	return n.colorOf() == RED
}

func (n *Node) uncle() *Node {
	if n == nil {
		return nil
	}
	return n.parent.sibling()
}

//sibling 返回兄弟节点
func (n *Node) sibling() *Node {
	if n.isLeftChild() {
		return n.parent.right
	}
	if n.isRightChild() {
		return n.parent.left
	}
	return nil
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

//Add 添加
func (t *Tree) Add(element int64) {
	//添加的是根节点
	if t.root == nil {
		t.root = &Node{element: element, color: BLACK}
		return
	}
	node := t.add(element)
	t.afterAdd(node)

}

func (t *Tree) afterAdd(node *Node) {
	parent := node.parent
	if parent == nil {
		black(node)
		t.root = node
		return
	}
	//如果父节点是黑色直接返回
	if parent.isBlack() {
		return
	}
	uncle := parent.sibling()
	grand := parent.parent
	//叔父节点是红色
	if uncle.isRed() {
		black(parent)
		black(uncle)
		//把祖父节点染成red当作新添加的节点
		t.afterAdd(red(grand))
		return
	}
	//叔父节点不是红色，要进行旋转
	if uncle.isBlack() {
		if parent.isLeftChild() {
			//Step 1. parent染成black，grand染成red
			red(grand)
			if node.isLeftChild() { //LL
				black(parent)
			} else { //LR
				black(node)
				t.rotateLeft(parent)
			}
			//Step 2. 将grand进行单旋
			t.rotateRight(grand)
		} else {
			red(grand)
			if node.isLeftChild() { //RL
				black(node)
				t.rotateRight(parent)

			} else { //RR
				//Step 1. parent染成black，grand染成red
				black(parent)
			}
			//Step 2. 将grand进行单旋
			t.rotateLeft(grand)
		}
	}
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
		color:   RED,
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
	}
	//更新child的parent
	if child != nil {
		child.parent = grand
	}
	grand.parent = parent
}

//Remove 移除
func (t *Tree) Remove(element int64) bool {
	node, err := t.remove(element)
	fmt.Println(node)
	if err == nil {
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
		t.afterRemove(node, replacement)
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
		//删除节点后处理
		t.afterRemove(node, nil)
	}

	return node, nil
}

func (t *Tree) afterRemove(n, r *Node) {
	//情况1. 删除的是红色节点
	if n.isRed() {
		return
	}
	//判断：用以替代的子节点是red
	if r.isRed() == true {
		black(r)
		return
	}
	//情况2. 删除的是黑色叶子节点
	parent := n.parent
	//1. 根节点
	if parent == nil {
		t.root = nil
		return
	}
	//2. 判断被删除的node是左还是右
	var left bool
	var sibling *Node
	if parent.left == nil || n.isLeftChild() {
		sibling = parent.right
		left = true
	} else {
		sibling = parent.left
		left = false
	}
	if left { //被删除的节点在左边，兄弟节点在右边
		//兄弟节点是红色
		if sibling.isRed() {
			//兄弟节点染成黑色，parent染成红色
			black(sibling)
			red(parent)
			//左旋
			t.rotateLeft(parent)
			//更换兄弟
			sibling = parent.right
		}
		//兄弟节点是黑色
		if sibling.left.isBlack() && sibling.right.isBlack() {
			//兄弟节点没有1个节点红色子节点，父节点要向下跟兄弟节点合并
			//判断parent是否黑色
			parentBlack := parent.isBlack()
			black(parent)
			red(sibling)
			if parentBlack {
				//下溢
				t.afterRemove(parent, nil)
			}
		} else {
			//兄弟节点至少有一个红的，向兄弟节点借元素

			//左边是黑色，兄弟要先旋转
			if sibling.right.isBlack() {
				t.rotateRight(sibling)
				sibling = parent.right
			}
			//给兄弟节点染色
			color(sibling, parent.colorOf())
			black(sibling.right)
			black(parent)
			t.rotateLeft(parent)
		}
	} else { //被删除的节点在右边，兄弟节点在左边
		//兄弟节点是红色
		if sibling.isRed() {
			//兄弟节点染成黑色，parent染成红色
			black(sibling)
			red(parent)
			//右旋
			t.rotateRight(parent)
			//更换兄弟
			sibling = parent.left
		}
		//兄弟节点是黑色
		if sibling.left.isBlack() && sibling.right.isBlack() {
			//兄弟节点没有1个节点红色子节点，父节点要向下跟兄弟节点合并
			//判断parent是否黑色
			parentBlack := parent.isBlack()
			black(parent)
			red(sibling)
			if parentBlack {
				//下溢
				t.afterRemove(parent, nil)
			}
		} else {
			//兄弟节点至少有一个红的，向兄弟节点借元素

			//左边是黑色，兄弟要先旋转
			if sibling.left.isBlack() {
				t.rotateLeft(sibling)
				sibling = parent.left
			}
			//给兄弟节点染色
			color(sibling, parent.colorOf())
			black(sibling.left)
			black(parent)
			t.rotateRight(parent)
		}
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
	fmt.Println(pre)
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

//LevelOrderTraversal 层序遍历
func (t *Tree) LevelOrderTraversal() []NodeValue {
	var list []NodeValue
	if t.root == nil {
		return list
	}
	queue := make([]*Node, 0)
	queue = append(queue, t.root)

	for len(queue) != 0 {
		tmp := make([]*Node, 0)
		for _, v := range queue {
			list = append(list, NodeValue{v.element, v.color})
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

//Clear 清除树
func (t *Tree) Clear() {
	t.root = nil
}
