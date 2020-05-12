package sort

//insertSort 插入排序
func insertSort(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		for j := i; j > 0; j-- {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			} else {
				break
			}
		}
	}
	return nums
}
