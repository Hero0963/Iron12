package binarySearch

func SearchMethod1(nums []int, target int) int {

	for i, v := range nums {
		if target == v {
			return i
		}

		if v > target {
			return -1
		}
	}

	return -1
}

func SearchMethod2(nums []int, target int) int {

	head := 0
	tail := len(nums) - 1
	middle := (head + tail) / 2

	for head <= tail {
		middle = (head + tail) / 2

		if nums[middle] == target {
			return middle
		} else if nums[middle] > target {
			tail = middle - 1
		} else {
			head = middle + 1
		}
	}

	return -1
}
