package pascalsTriangleII

func GetRowMethod1(rowIndex int) []int {
	ans := make([]int, 1, rowIndex+1)
	ans[0] = 1
	if rowIndex == 0 {
		return ans
	}

	for i := 0; i < rowIndex; i++ {
		ans = append(ans, 1)
		for j := len(ans) - 2; j > 0; j-- {
			ans[j] += ans[j-1]
		}
	}

	return ans
}
