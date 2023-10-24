//哈希
func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i, v := range nums {
		p, ok := m[target-v]
		if ok == true {
			return []int{p, i}
		} else {
			m[v] = i
		}
	}
	return []int{}
}

//暴力
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}