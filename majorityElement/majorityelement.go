package majorityElement

func MajorityElementMethod1(nums []int) int {

	m := make(map[int]int)
	half := len(nums) / 2

	for _, v := range nums {
		m[v]++
		if m[v] > half {
			return v
		}

	}

	return -1
}

func MajorityElementMethod2(nums []int) int {
	candidate := nums[0]
	count := 0

	for _, v := range nums {

		if count == 0 {
			candidate = v
			count++
			continue
		}

		if candidate == v {
			count++
		} else {
			count--
		}

	}

	return candidate
}
