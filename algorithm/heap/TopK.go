package heap

//Top K问题
//解题思路
//法一：用快排；复杂度O(n log n)
//法二：用小顶堆排序 O(n log k)
//Step 1. 创建一个小顶堆
//Step 2. 扫描前k个元素入堆
//Step 2. 再将k+1以后的元素，和堆顶的元素进行比较；如果大于堆顶，替换堆顶元素

func TopK(list []int, k int) []int {
	heap := NewMinHeap()
	for i := 0; i < len(list); i++ {
		if i < k {
			heap.Add(list[i])
		} else {
			if list[i] > heap.element[0] {
				heap.Replace(list[i])
			}
		}
	}
	return heap.element
}
