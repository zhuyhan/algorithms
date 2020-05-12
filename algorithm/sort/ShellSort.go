package sort

//shellSort 希尔排序
func shellSort(nums []int) []int {
	//1.根据数组长度确定增长量h
	var h int = 1
	for h < len(nums)/2 {
		h = 2*h + 1
	}
	//2.希尔排序
	for h >= 1 {
		for i := h; i < len(nums); i++ {
			for j := i; j >= h; j -= h {
				//2.1 找到待插入的元素
				if nums[j-h] > nums[j] {
					//2.2 把待插入的元素插入插入有序数列中
					nums[j-h], nums[j] = nums[j], nums[j-h]
				}
			}
		}
		//减小h的值
		h = h / 2
	}
	return nums
}
