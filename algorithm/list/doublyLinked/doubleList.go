package doublyLinked

import (
	//"container/list"
	"errors"
)

type Element struct {
	next, prev *Element
	list       *List
	Value      interface{}
}
type List struct {
	root Element
	len  int
}

func New() *List {
	return new(List).Init()
}

func (l *List) Init() *List {
	l.root.prev = &l.root
	l.root.next = &l.root
	l.len = 0
	return l
}

//Front 第一个结点
func (l *List) Front() *Element {
	if l.len == 0 {
		return nil
	}
	return l.root.next
}

//Black 尾结点
func (l *List) Black() *Element {
	if l.len == 0 {
		return nil
	}
	return l.root.prev
}

//Node 获取index位置对应的结点
func (l *List) Node(index int) *Element {
	var node *Element
	if index < 0 || index >= l.len {
		return node
	}
	if index < (l.len >> 1) {
		for i := 0; i < index; i++ {
			node = l.root.next
		}
	} else {
		for i := l.len - 1; i > index; i-- {
			node = l.root.prev
		}
	}
	return node
}

func (l *List) Add(index int, element interface{}) (*List, error) {
	if index < 0 || index > l.len {
		return l, errors.New("index is error")
	}
	newNode := &Element{
		next:  nil,
		prev:  nil,
		list:  l,
		Value: element,
	}

	if index == l.len {
		l.root.next = newNode
		l.root.prev = newNode
	} else {
		nextNode := l.Node(index)
		preNode := nextNode.prev
		newNode.prev = preNode
		newNode.next = nextNode
		if preNode == nil {
			l.root.prev = newNode
		} else {
			preNode.next = newNode
		}
		if nextNode == nil {
			l.root.next = newNode
		} else {
			nextNode.prev = newNode
		}

	}

	l.len++
	return l, nil
}
