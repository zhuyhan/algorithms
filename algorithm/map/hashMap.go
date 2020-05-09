package _map

import (
	tree "algorithms/algorithm/binary_tree/red_black_tree"
)

const (
	DefaultCapacity = 1 << 4
)

type HashMap struct {
	size  int
	table []*tree.Tree
}

func New(size int) *HashMap {
	//if size <= DefaultCapacity {
	//	size = DefaultCapacity
	//}
	return &HashMap{
		size:  0,
		table: nil,
	}
}

func (h *HashMap) clear() {
	if h.size == 0 {
		return
	}
	h.size = 0
	len := len(h.table)
	for i := 0; i < len; i++ {
		h.table[i].Clear()
	}
	return
}

func hashCode(str string) int {
	var hashCode int
	for _, v := range str {
		hashCode = hashCode*31 + int(v)
	}
	return hashCode
}
