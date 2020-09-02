package twoSum

func TwoSum(nums []int, target int) []int {

	m := make(map[int]int)

	for i, v := range nums {

		j, exist := m[target-v]

		if exist {
			return []int{i, j}
		}

		m[v] = i
	}

	return []int{-1, -1}
}
