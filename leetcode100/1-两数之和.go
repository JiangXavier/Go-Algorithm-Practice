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