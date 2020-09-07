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

func SearchMethod3(nums []int, target int) int {
	return helper(nums, target, 0, len(nums)-1)

}

func helper(nums []int, target int, head int, tail int) int {
	if head >= tail {
		if nums[head] == target {
			return head
		}
		return -1
	}
	if head < 0 {
		return -1
	}
	if tail == len(nums) {
		return -1
	}

	mid := head + (tail-head)/2
	switch {
	case nums[mid] > target:
		return helper(nums, target, head, mid-1)
	case nums[mid] < target:
		return helper(nums, target, mid+1, tail)
	default:
		return mid
	}
}
