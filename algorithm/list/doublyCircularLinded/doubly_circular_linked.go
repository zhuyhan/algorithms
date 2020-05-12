package doublyCircularLinded

import (
	"errors"
)

//双向循环链表的数据结构实现

type Node struct {
	value      int
	next, prev *Node
}

type List struct {
	first *Node //指向头节点
	last  *Node
	len   int
}

func New() *List {
	return new(List).Init()
}

func NewNode(prev, next *Node, element int) *Node {
	return &Node{
		value: element,
		next:  next,
		prev:  prev,
	}
}

func (l *List) Init() *List {
	l.first = nil
	l.last = nil
	l.len = 0
	return l
}

func (l *List) Add(element int) {
	node := &Node{value: element}
	if l.len == 0 {
		l.first = node
		node.prev = node
		node.next = node
	} else {
		lastNode := l.last   //最后一个节点
		firstNode := l.first //第一个节点
		firstNode.prev = node
		lastNode.next = node
		node.prev = lastNode
		node.next = firstNode
	}
	l.last = node
	l.len++
	return
}

func (l *List) Insert(index, element int) error {
	if err := l.rangeCheck(index); err != nil {
		return err
	}

	if index == l.len {
		node := NewNode(l.last, l.first, element)
		oldLastNode := l.last
		if oldLastNode == nil {
			//当链表为空时
			l.first = node
			l.first.prev = l.first
			l.first.next = l.last
		} else {
			oldLastNode.next = node
			l.first.prev = node
		}
	} else {
		nextNode, err := l.Node(index)
		if err != nil {
			return err
		}
		prevNode := nextNode.prev
		node := NewNode(prevNode, nextNode, element)
		nextNode.prev = node
		prevNode.next = node
		if index == 0 {
			//当index == 0时
			l.first = node
		}

	}
	l.len++
	return nil
}

func (l *List) Remove(index int) error {
	if index < 0 || index > l.len-1 {
		return errors.New("index is error")
	}
	if l.len <= 1 {
		l.Init()
		return nil
	}
	node, _ := l.Node(index)
	prev := node.prev
	next := node.next
	prev.next = next
	next.prev = prev
	if node == l.first {
		l.first = next
	}
	if node == l.last {
		l.last = prev
	}
	l.len--
	return nil
}

func (l *List) Node(index int) (*Node, error) {
	if err := l.rangeCheck(index); err != nil {
		return nil, err
	}
	node := l.first
	for i := 0; i < index; i++ {
		node = node.next
	}
	return node, nil
}

//rangeCheck index的范围确认
func (l *List) rangeCheck(index int) error {
	if index < 0 || index > l.len {
		return errors.New("index is error")
	}
	return nil
}

func (l *List) GetAll() []int {
	node := l.first
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
