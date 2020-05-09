package trie

import "fmt"

//字典树
type Trie struct {
	root *Node
}

type Node struct {
	last     bool //判断是否为单词结尾
	parent   *Node
	children map[int32]*Node
}

func NewTrie() *Trie {
	return &Trie{root: NewNode()}
}

func NewNode() *Node {
	return &Node{
		children: make(map[int32]*Node),
	}
}

//add 添加单词
func (t *Trie) Add(word string) {
	node := t.root
	for _, v := range word {
		if _, ok := node.children[v]; ok {
			//已存在
			node = node.children[v]
		} else {
			//不存在,添加子节点
			childNode := &Node{
				children: map[int32]*Node{},
				parent:   node,
			}
			node.children[v] = childNode
			node = childNode
		}
	}
	node.last = true
}

//Search 查找单词(完整的单词)
func (t *Trie) Search(word string) bool {
	node := t.root
	for _, v := range word {
		fmt.Println(node)
		if _, ok := node.children[v]; ok {
			node = node.children[v]
		} else {
			return false
		}
	}
	if node.last == false {
		return false
	}
	return true
}

//Remove 删除
//Step 1. 找到单词（last 必须为 true）；修改last为false
//Step 2. 如果最后一个节点无子节点，删除该节点（向上删除）
func (t *Trie) Remove(word string) bool {
	node := t.root
	for _, v := range word {
		if _, ok := node.children[v]; !ok {
			return false
		} else {
			node = node.children[v]
		}
	}
	if node.last == false {
		return false
	}
	node.last = false
	index := len(word) - 1
	for len(node.children) == 0 {
		delete(node.children, int32(word[index]))
		node = node.parent
		index--
	}
	return true
}

func (t *Trie) StartWith(prefix string) bool {
	node := t.root
	for _, v := range prefix {
		if _, ok := node.children[v]; ok {
			node = node.children[v]
		} else {
			return false
		}
	}
	return true
}
