package list

func twoSum(nums []int, target int) []int {
	if len(nums) <= 2 {
		return []int{-1, -1}
	}
	numsMap := make(map[int]int)
	for t1, v := range nums {
		if t2, ok := numsMap[target-v]; ok {
			return []int{t2, t1}
		}
		numsMap[v] = t1
	}
	return []int{-1, -1}
}
