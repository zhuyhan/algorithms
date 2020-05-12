package sort

//基数排序
//非常适合用于整数排序（特别是非负整数），因此设定元素>=0
//执行流程：从低位到高位排序

func radixSort(nums []int) []int {
	//	找到最大值
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if max < nums[i] {
			max = nums[i]
		}
	}
	//位数-593
	//个位数 nums[i] / 1 % 10 = 3
	//十位数 nums[i] / 10 % 10 = 9
	//百位数 nums[i] / 100 % 10 = 5
	//百位数 nums[i] / 1000 % 10 = 0
	for divider := 1; divider < max; divider *= 10 {
		nums = countingForRadixSort(nums, divider)
	}
	return nums
}

//countingForRadixSort 计数排序
func countingForRadixSort(nums []int, divider int) []int {
	//开辟内存空间，存储次数
	counts := make([]int, 10)
	//统计每个整数出现的次数
	for i := 0; i < len(nums); i++ {
		counts[nums[i]/divider%10]++
	}
	//累加次数
	for i := 1; i < len(counts); i++ {
		counts[i] += counts[i-1]
	}
	//从后往前遍历元素，将它放入有序数组
	result := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		counts[nums[i]/divider%10] -= 1
		result[counts[nums[i]/divider%10]] = nums[i]
	}
	return result
}
