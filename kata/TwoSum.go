package kata

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for index, num := range nums {
		if v, found := m[target-num]; found {
			return []int{v, index}
		}

		m[num] = index
	}
	return nil
}
