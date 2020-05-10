package singleLinked

import (
	"fmt"
)

//单向链表的数据结构(含虚拟头结点)

//结点
type Node struct {
	data interface{}
	next *Node
}

//单向链表
type List struct {
	head *Node
}

func NewLinkList() *List {
	head := &Node{
		data: 0,
		next: nil,
	}
	return &List{
		head: head,
	}
}

//Append 尾部添加数据
func (l *List) Add(elem interface{}) bool {
	//构造要插入的结点
	node := &Node{
		data: elem,
		next: nil,
	}
	//当前循环到结点
	currentNode := l.head
	for currentNode.next != nil {
		currentNode = currentNode.next
	}
	currentNode.next = node
	return true
}

//Insert 往新结点插入
func (l *List) Insert(index int, elem interface{}) bool {
	if index < 0 || index > l.Length() {
		fmt.Println("error:index=", index)
		return false
	}
	node := &Node{
		data: elem,
		next: nil,
	}
	currentNode := l.head
	for i := 0; i < index; i++ {
		currentNode = currentNode.next
	}
	node.next = currentNode.next
	currentNode.next = node
	return true

}

//Length 长度
func (l *List) Length() int {
	currentNode := l.head
	length := 0
	for currentNode.next != nil {
		currentNode = currentNode.next
		length++
	}
	return length
}

//Delete 指定删除索引
func (l *List) Delete(index int) bool {
	//判断index是否合法
	if index < 0 || index > l.Length()-1 {
		fmt.Println("删除位序不正确")
		return false
	}
	currentNode := l.head
	for i := 0; i < index; i++ {
		currentNode = currentNode.next
	}
	currentNode.next = currentNode.next.next
	return true
}

//GetData 通过索引，获取元素值
func (l *List) GetData(index int) interface{} {
	if index < 0 || index > l.Length()-1 {
		fmt.Println("位序不正确")
		return false
	}
	currentNode := l.head
	for i := 0; i <= index; i++ {
		currentNode = currentNode.next
	}
	return currentNode.data
}

func (l *List) GetAll() []interface{} {
	var list []interface{}
	currentNode := l.head
	for currentNode.next != nil {
		currentNode = currentNode.next
		list = append(list, currentNode.data)
	}
	return list
}
