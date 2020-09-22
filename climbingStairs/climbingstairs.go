package climbingStairs

func ClimbStairsMethod1(n int) int {

	if n <= 2 {
		return n
	}

	first := 1
	second := 2
	sum := 0

	for i := 3; i <= n; i++ {
		sum = first + second
		first = second
		second = sum
	}

	return sum
}
