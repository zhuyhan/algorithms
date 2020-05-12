package sort

//selectionSort2 选择排序
func selectionSort2(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(nums); j++ {
			if nums[minIndex] > nums[j] {
				minIndex = j
			}
		}
		if minIndex != i {
			nums[minIndex], nums[i] = nums[i], nums[minIndex]
		}
	}
	return nums
}

func selectionSort(nums []int) []int {
	for i := len(nums) - 1; i > 0; i-- {
		pos := i
		for j := 0; j < i; j++ {
			if nums[pos] < nums[j] {
				pos = j
			}
		}
		if pos != i {
			swap(nums, pos, i)
		}
	}
	return nums
}
