package heap

//二叉堆-最大堆

type MaxHeap struct {
	element []int
}

func Init() *MaxHeap {
	return &MaxHeap{
		element: make([]int, 0),
	}
}

//Add 添加元素到堆末尾，然后进行上浮操作 O(log n)
func (b *MaxHeap) Add(element int) {
	b.element = append(b.element, element)
	index := len(b.element) - 1
	b.shiftUp(index)
	return
}

//Remove 删除堆顶元素 O(log n)
//将堆底元素与堆顶替换，删除；然后下沉
func (b *MaxHeap) Remove() {
	if len(b.element) == 0 {
		return
	}
	lastIndex := len(b.element) - 1
	//替换堆顶
	b.element[0] = b.element[lastIndex]
	//删除最后一个元素
	b.element = b.element[:lastIndex]
	b.shiftDown(0)
}

//Replace 删除堆顶元素，同时新增一个元素
func (b *MaxHeap) Replace(element int) {
	if len(b.element) == 0 {
		b.element = append(b.element, element)
		return
	}
	//替换堆顶
	b.element[0] = element
	b.shiftDown(0)
}

//CreatHeap 批量建堆
//法一：自上而下的上滤O(n log n)
func Heapify1(list []int) []int {
	heap := Init()
	heap.element = list
	for k, _ := range list {
		heap.shiftUp(k)
	}
	return heap.element
}

//法一：自下而上️的下滤 （推荐	O(n)）
func Heapify2(list []int) []int {
	heap := Init()
	heap.element = list
	half := len(heap.element) >> 1 //第一个非叶子节点
	for i := half - 1; i >= 0; i-- {
		heap.shiftDown(i)
	}
	return heap.element
}

func (b *MaxHeap) shiftUp(index int) {
	element := b.element[index]
	for index > 0 {
		pindex := (index - 1) >> 1
		p := b.element[pindex] //父节点的值
		if p >= element {
			break
		}
		b.element[index] = p
		index = pindex
	}
	b.element[index] = element
}
func (b *MaxHeap) shiftDown(index int) {
	element := b.element[index]
	len := len(b.element)
	half := len >> 1
	//必须保证index是非叶子节点
	for index < half {
		//1.只有左子节点
		//2.index有两个节点

		//默认为左子节点跟他比较
		childIndex := (index << 1) + 1
		child := b.element[childIndex]

		//判断是否有右子节点，并大于左子树
		rightIndex := childIndex + 1
		if rightIndex < len && b.element[rightIndex] > child {
			childIndex = rightIndex
			child = b.element[childIndex]
		}
		if element >= child {
			break
		}
		//将子节点放在index位置
		b.element[index] = child
		index = childIndex
	}
	b.element[index] = element

}

func (b *MaxHeap) isEmpty() bool {
	return len(b.element) == 0
}

func (b *MaxHeap) clear() {
	b.element = nil
}
