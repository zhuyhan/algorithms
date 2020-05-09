package set

func intersection(nums1 []int, nums2 []int) []int {
	arr := make(map[int]int, 0)
	var list []int
	for _, v := range nums1 {
		if _, ok := arr[v]; !ok {
			arr[v] = 1
		}
	}
	for _, key := range nums2 {
		if v, ok := arr[key]; ok {
			if v == 1 {
				list = append(list, key)
			}
			arr[key] += 2
		} else {
			arr[key] = 2
		}
	}
	return list
}
