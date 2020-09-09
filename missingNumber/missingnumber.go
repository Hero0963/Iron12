package missingNumber

import "sort"

func MissingNumberMethod1(nums []int) int {
	m := make(map[int]bool)
	for _, v := range nums {
		m[v] = true
	}

	for i := 0; i < len(nums); i++ {
		if m[i] == false {
			return i
		}
	}

	return len(nums)
}

func MissingNumberMethod2(nums []int) int {
	sum := 0
	sum = ((len(nums) + 1) * len(nums)) / 2

	for _, v := range nums {
		sum -= v
	}

	return sum
}

func MissingNumberMethod3(nums []int) int {
	ans := 0
	ans ^= len(nums)

	for i, v := range nums {
		ans ^= i ^ v
	}

	return ans
}

func MissingNumberMethod4(nums []int) int {
	sort.Ints(nums)

	head := 0
	tail := len(nums)
	mid := 0

	for head < tail {
		mid = head + (tail-head)/2

		if nums[mid] > mid {
			tail = mid
		} else {
			head = mid + 1
		}
	}

	return tail
}
