package hash_table

import (
	list "algorithms/algorithm/list/doublyLinked"
	"errors"
)

const (
	DefaultCapacity = 1 << 4
)

type item struct {
	key   string
	value interface{}
}

type HashTable struct {
	size     int
	capacity int
	table    map[int]*list.List
}

func New(cap int) *HashTable {
	if cap <= 0 {
		cap = DefaultCapacity
	}
	return &HashTable{
		size:     0,
		capacity: cap,
		table:    make(map[int]*list.List, cap),
	}
}

//Get 获取哈希表的值
func (h *HashTable) Get(key string) (interface{}, error) {
	index := h.hash(key)
	item, err := h.find(index, key)
	if item == nil {
		return "", errors.New("not found")
	}
	return item.value, err
}

//Put 写入哈希表
func (h *HashTable) Put(key, value string) {
	index := h.hash(key)
	//如果哈希表没有有该数据，创建一个新的链表
	if h.table[index] == nil {
		h.table[index] = list.New()
	}
	item := &item{key: key, value: value}
	node, err := h.find(index, key)
	if err != nil {
		h.table[index].Add(item)
		h.size++
	} else {
		//已存在该key，覆盖value
		node.value = value
	}
}

//Remove 移除元素
func (h *HashTable) Remove(key string) error {
	index := h.hash(key)
	linked, ok := h.table[index]
	if !ok {
		return nil
	}
	var val *item
	linked.Each(func(node list.Node) {
		if node.Value.(*item).key == key {
			val = node.Value.(*item)
		}
	})
	if val == nil {
		return nil
	}
	h.size--
	return linked.Del(val)
}

func (h *HashTable) find(index int, key string) (*item, error) {
	var result *item
	linked, ok := h.table[index]
	if !ok {
		return nil, errors.New("not found")
	}

	linked.Each(func(node list.Node) {
		if node.Value.(*item).key == key {
			result = node.Value.(*item)
		}
	})
	if result == nil {
		return nil, errors.New("not found")
	}
	return result, nil
}

func (h *HashTable) hash(str string) int {
	return hashCode(str) % h.capacity
}

//hashCode 哈希函数
func hashCode(str string) int {
	var hashCode int
	for _, v := range str {
		hashCode = hashCode*31 + int(v)
	}
	return hashCode
}
