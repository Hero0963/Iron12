package singleNumber

func SingleNumberMethod1(nums []int) int {
	m := make(map[int]bool)

	for _, v := range nums {

		_, exist := m[v]

		if exist {
			delete(m, v)
		} else {
			m[v] = true
		}
	}

	for key := range m {
		if m[key] == true {
			return key
		}
	}

	return -1
}

func SingleNumberMethod2(nums []int) int {
	flag := 0
	for _, v := range nums {
		flag ^= v
	}
	return flag
}
