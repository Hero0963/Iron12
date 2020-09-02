package twoSum

func TwoSumMethod1(nums []int, target int) []int {

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{-1, -1}
}

func TwoSumMethod2(nums []int, target int) []int {

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
