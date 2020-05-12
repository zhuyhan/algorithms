package sort

//执行流程
//Step 1. 对序列进行原地建堆
//Step 2. 交换堆顶元素和尾元素
//Step 3. 堆元素的数量减1，对0位置进行一次shiftDown操作
//Step 4. 重复Step 2和Step 3，直到堆元素的数量为1

func heapSort(nums []int) []int {
	//原地建堆
	for i := (len(nums) >> 1) - 1; i >= 0; i-- {
		shiftDown(nums, i, len(nums))
	}
	//交换堆顶元素和尾元素
	for size := len(nums) - 1; size > 0; size-- {
		swap(nums, 0, size)
		shiftDown(nums, 0, size)
	}
	return nums
}

func shiftDown(nums []int, index, length int) {
	element := nums[index]
	half := length >> 1
	for index < half {
		//默认孩子节点为左子节点
		childIndex := index<<1 + 1
		child := nums[childIndex]
		//判断是否有右子节点,选择值大的子节点替换
		rightIndex := childIndex + 1
		if rightIndex < length && child < nums[rightIndex] {
			childIndex = rightIndex
			child = nums[rightIndex]
		}
		if element >= child {
			break
		}
		nums[index] = child
		index = childIndex
	}
	nums[index] = element
}
