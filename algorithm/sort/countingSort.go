package sort

//计数排序
//1. 只能满足一定范围内的整数进行排序
//2. 空间复杂度O(n+k)
//3. 最好、最坏、平均时间复杂度O(n+k)
//4. 属于稳定排序
func countingSort(nums []int) []int {
	// 找到最大值、最小值
	min, max := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if min > nums[i] {
			min = nums[i]
		}
		if max < nums[i] {
			max = nums[i]
		}
	}

	//开辟内存空间，存储每个整数出现的次数
	counts := make([]int, max-min+1)
	for _, i := range nums {
		counts[i-min]++
	}
	//累加次数
	for i := 1; i < len(counts); i++ {
		counts[i] += counts[i-1]
	}
	//从后往前遍历元素，将它放入有序数组
	result := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		counts[nums[i]-min] -= 1
		result[counts[nums[i]-min]] = nums[i]

	}
	return result
}
