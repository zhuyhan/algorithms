package sort

func testMergeSort(nums []int) []int {
	mergeSort(nums, 0, len(nums)-1)
	return nums
}

//mergeSort 归并排序
//l为最小索引，r为最大索引
func mergeSort(nums []int, l, r int) {
	if l < r {
		mid := (l + r) / 2
		//对每组数据进行排序
		mergeSort(nums, l, mid)
		mergeSort(nums, mid+1, r)
		merge(nums, l, mid, r)
	}
}

func merge(nums []int, l, mid, r int) {
	//左右两个子数组
	n1, n2 := mid-l+1, r-mid
	left := make([]int, n1)
	right := make([]int, n2)
	copy(left, nums[l:mid+1])
	copy(right, nums[mid+1:r+1])
	//将左右子数组排序合并到原数组中
	i, k, j := 0, l, 0
	for ; i < n1 && j < n2; k++ {
		if left[i] > right[j] {
			nums[k] = right[j]
			j++
		} else {
			nums[k] = left[i]
			i++
		}
	}
	for ; i < n1; i++ {
		nums[k] = left[i]
		k++
	}
	for ; j < n2; j++ {
		nums[k] = right[j]
		k++
	}
}
