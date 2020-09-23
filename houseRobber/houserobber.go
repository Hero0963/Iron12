package houseRobber

func RobMethod1(nums []int) int {

	var pre2, pre1, now int
	for _, v := range nums {
		now = max(pre2+v, pre1)
		pre2 = pre1
		pre1 = now
	}

	return now
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
