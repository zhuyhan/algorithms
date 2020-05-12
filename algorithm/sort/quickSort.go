package sort

func testQuickSort(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

//快速排序
func quickSort(nums []int, l, r int) {
	if l < r {
		var (
			i     = l
			j     = r
			pivot = nums[l]
		)
		//对l到r索引处的元素进行分组
		for i < j {
			for i < j && nums[j] > pivot {
				j--
			}
			for i < j && nums[i] <= pivot {
				i++
			}
			if i < j {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
		nums[i], nums[l] = nums[l], nums[i]
		//让左子组有序
		quickSort(nums, l, i-1)
		//让右子组有序
		quickSort(nums, i+1, r)
	}
}
