package heap

//小顶堆的实现

type MinHeap struct {
	element []int
}

func NewMinHeap() *MinHeap {
	return &MinHeap{
		element: make([]int, 0),
	}
}

//Add 添加元素
func (m *MinHeap) Add(element int) {
	m.element = append(m.element, element)
	//长度=0，直接添加为根节点
	if len(m.element) == 0 {
		return
	}
	index := len(m.element) - 1
	//元素上浮
	m.shiftUp(index)
	return
}

//shiftUp 上浮
func (m *MinHeap) shiftUp(index int) {
	element := m.element[index]
	for index > 0 {
		pindex := (index - 1) >> 1  //父节点索引
		parent := m.element[pindex] //父节点
		if parent <= index {
			break
		}
		m.element[index] = parent
		index = pindex
	}
	m.element[index] = element
	return
}

//Remove 移除堆顶元素
func (m *MinHeap) Remove(element int) {
	if len(m.element) == 0 {
		return
	}
	//用最后一个元素，替换堆顶元素；
	lastIndex := len(m.element) - 1
	m.element[0] = m.element[lastIndex]
	m.element = m.element[:lastIndex]
	//然后删除最后一个元素，进行下沉操作
	m.shiftDown(0)
}
func (m *MinHeap) shiftDown(index int) {
	element := m.element[index]
	half := len(m.element) >> 1
	for index < half {
		//默认child节点为左子节点
		childIndex := (index << 1) + 1
		child := m.element[childIndex]
		//判断右右子树，并且右子树的值小于左子树的值
		rightIndex := childIndex + 1
		if rightIndex < len(m.element) && m.element[rightIndex] < child {
			childIndex = rightIndex
			child = m.element[rightIndex]
		}
		//判断大小，下沉
		if element <= child {
			break
		}
		m.element[index] = child
		index = childIndex
	}
	m.element[index] = element
	return
}

//Replace 替换操作
func (m *MinHeap) Replace(element int) {
	if len(m.element) == 0 {
		m.element = append(m.element, element)
	}
	m.element[0] = element
	m.shiftDown(0)
}

//Heapify 自下而上的下沉
func Heapify(list []int) []int {
	heap := &MinHeap{element: list}
	half := len(list) >> 1
	for i := half - 1; i >= 0; i-- {
		heap.shiftDown(i)
	}
	return heap.element
}
