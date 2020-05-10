package doublyLinked

import (
	//"container/list"
	"errors"
)

//双向链表

type Node struct {
	next, prev *Node
	Value      interface{}
}

type List struct {
	first *Node //指向头节点
	last  *Node
	len   int
}

func New() *List {
	return new(List).Init()
}

func (l *List) Init() *List {
	l.first = nil
	l.last = nil
	l.len = 0
	return l
}

//Front 第一个结点
func (l *List) Front() *Node {
	if l.len == 0 {
		return nil
	}
	return l.first
}

//Black 尾结点
func (l *List) Black() *Node {
	if l.len == 0 {
		return nil
	}
	return l.last
}

//Node 获取index位置对应的结点
func (l *List) Node(index int) *Node {
	if index < 0 || index >= l.len {
		return nil
	}
	var node *Node
	if index < (l.len >> 1) {
		node = l.first
		//索引小于长度的一半
		for i := 0; i < index; i++ {
			node = node.next
		}
	} else {
		node = l.last
		for i := l.len - 1; i > index; i-- {
			node = node.prev
		}
	}
	return node
}

//Add 往最后添加元素
func (l *List) Add(element interface{}) {
	node := &Node{Value: element}
	if l.first == nil {
		l.first = node
	} else {
		l.last.next = node
		node.prev = l.last
	}
	l.last = node
	l.len++
	return
}

//Insert 往索引位置插入一个元素
func (l *List) Insert(index int, element interface{}) (*List, error) {
	if index < 0 || index > l.len {
		return l, errors.New("index is error")
	}
	node := &Node{Value: element}
	if index == l.len { // 往最后添加元素
		last := l.last
		if last == nil { //链表为空
			l.first = node
		} else {
			node.prev = last
			last.next = node
		}
		l.last = node
	} else {
		nextNode := l.Node(index)
		if nextNode == nil {
			return l, errors.New("index is nil")
		}
		prevNode := nextNode.prev
		node.prev = prevNode
		node.next = nextNode
		if prevNode == nil {
			//前驱节点为空，那么index=0
			l.first = node
		} else {
			prevNode.next = node
		}
	}
	l.len++
	return l, nil
}

func (l *List) Remove(index int) error {
	if l.len == 0 {
		return errors.New("list is empty")
	}
	if index < 0 || index > l.len {
		return errors.New("index is error")
	}
	nextNode := l.Node(index)
	prevNode := nextNode.prev
	if prevNode == nil {
		//	index == 0
		l.first = nextNode
	} else {
		prevNode.next = nextNode
	}
	if nextNode == nil {
		//index == l.len-1
		l.last = prevNode
	} else {
		nextNode.prev = prevNode
	}
	l.len--
	return nil
}

//Clear 清空链表
func (l *List) Clear() {
	l.Init()
}

//获取链表所有元素
func (l *List) GetAll() []interface{} {
	node := l.first
	result := make([]interface{}, 0)
	for node != nil {
		result = append(result, node.Value)
		node = node.next
	}
	return result
}
