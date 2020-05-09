package array_list

//动态数组数据结构实现

import (
	"errors"
	"fmt"
)

const (
	DefaultCapacity = 10
	ElementNotFound = -1
)

type Array struct {
	size    int   //元素数量
	element []int //所有元素
}

//NewArray 初始化数据的容量为capacity
func NewArray(capacity int) *Array {
	array := new(Array)
	if capacity <= 0 {
		capacity = DefaultCapacity
	}
	array.element = make([]int, capacity)
	return array
}

//IsEmpty 是否为空
func (a *Array) IsEmpty() bool {
	return a.size == 0
}

//Get 通过索引获取元素
func (a *Array) Get(index int) (int, error) {
	if index < 0 || index > a.size {
		return 0, errors.New(fmt.Sprintln("Index:", index, "; Size:", a.size))
	}
	return a.element[index], nil
}

//Set 设置index位置的元素
func (a *Array) Set(index int, element int) bool {
	if index < 0 || index > a.size {
		return false
	}
	a.element = append(a.element[0:index], append([]int{element}, a.element[index+1:]...)...)
	return true
}

//Add 插入一个元素
func (a *Array) Add(element int) bool {
	if a.size >= cap(a.element) {
		return false
	}
	a.element[a.size] = element
	a.size++
	return true
}

//IndexOf 查看元素第一次出现的索引值，若无返回-1
func (a *Array) IndexOf(element int) int {
	for i := 0; i < a.size; i++ {
		if a.element[i] == element {
			return i
		}
	}
	return ElementNotFound
}

//Contains 判断是否包含某个元素
func (a *Array) Contains(element int) bool {
	return a.IndexOf(element) != ElementNotFound
}

//Remove 删除索引所在元素
func (a *Array) Remove(index int) bool {
	if index < 0 || index > a.size {
		return false
	}
	a.element = append(a.element[:index], a.element[index+1:]...)
	a.size--
	return true
}

//Clear 清除数据
func (a *Array) Clear() {
	a.size = 0
}

//GetAll 获取所有数据
func (a *Array) GetAll() []int {
	return a.element[:a.size]
}
