package sort

//bubbleSort 冒泡排序
func bubbleSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}

//selectionSort 选择排序
func selectionSort(nums []int) []int {
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
