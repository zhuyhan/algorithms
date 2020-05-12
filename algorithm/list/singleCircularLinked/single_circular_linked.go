package singleCircularLinked

import (
	"errors"
	"fmt"
)

type Node struct {
	next  *Node
	value int
}

type List struct {
	root *Node
	len  int
}

func New() *List {
	return new(List).Init()
}

func (l *List) Init() *List {
	l.root = nil
	l.len = 0
	return l
}

func (l *List) Add(element int) {
	node := &Node{value: element}
	if l.root == nil {
		l.root = node
		node.next = node
	} else {
		prevNode, _ := l.Node(l.len - 1)
		node.next = prevNode.next
		prevNode.next = node
	}
	l.len++
	return
}

func (l *List) Insert(index, element int) error {
	if index < 0 || index > l.len {
		return errors.New("index is error")
	}
	node := &Node{value: element}
	if index == 0 {
		//拿到最后一个节点,将其next指向第一个节点
		if l.len == 0 {
			l.root = node
			node.next = node
		} else {
			lastNode, _ := l.Node(l.len - 1)
			lastNode.next = node
			node.next = l.root
		}
		l.root = node
	} else {
		prevNode, _ := l.Node(index - 1)
		node.next = prevNode.next
		prevNode.next = node
	}
	l.len++
	return nil
}

//Remove 删除节点
func (l *List) Remove(index int) error {
	if index < 0 || index >= l.len {
		return errors.New("index is error")
	}
	if index == 0 {
		//删除的是第一个节点
		if l.len == 1 {
			l.Clear()
		}
		lastNode, _ := l.Node(l.len - 1)
		lastNode.next = l.root.next
		l.root = l.root.next
	} else {

		prevNode, _ := l.Node(index - 1)
		fmt.Println(prevNode)
		node := prevNode.next
		prevNode.next = node.next
	}
	l.len--
	return nil
}

func (l *List) Node(index int) (*Node, error) {
	if index < 0 || index > l.len {
		return nil, errors.New("index is error")
	}
	node := l.root
	for i := 0; i < index; i++ {
		node = node.next
	}
	return node, nil
}

func (l *List) Clear() *List {
	return l.Init()
}

func (l *List) GetAll() []int {
	node := l.root
	result := make([]int, 0)
	for i := 0; i < l.len; i++ {
		if node == nil {
			break
		}
		result = append(result, node.value)
		node = node.next
	}
	return result
}
