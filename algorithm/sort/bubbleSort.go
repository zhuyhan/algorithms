package sort

//bubbleSort 冒泡排序
func bubbleSort(nums []int) []int {
	for i := len(nums) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if nums[j] > nums[j+1] {
				//nums[j], nums[j+1] = nums[j+1], nums[j]
				swap(nums, j, j+1)
			}
		}
	}
	return nums
}
