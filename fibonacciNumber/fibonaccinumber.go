package fibonacciNumber

import (
	"fmt"
	"math"
)

func FibMethod1(N int) int {
	ans := 0
	c := (1.0 / math.Sqrt(5))
	first := 1.0
	second := 1.0

	for i := 0; i < N; i++ {
		first *= (1 + math.Sqrt(5)) / 2.0
		second *= (1 - math.Sqrt(5)) / 2.0
	}

	ans = int(c * (first - second))
	return ans
}

func FibMethod2(N int) int {
	if N <= 1 {
		return N
	}

	return FibMethod2(N-1) + FibMethod2(N-2)

}

func FibMethod3(N int) int {
	if N <= 1 {
		return N
	}

	first := 0
	second := 1
	sum := 0

	for i := 2; i <= N; i++ {
		sum = first + second
		first = second
		second = sum
	}

	return sum
}

func FibMethod4(N int) int {

	fibNums := []int{0, 1}

	if N <= 1 {
		return N
	}

	first := 0
	second := 1
	sum := 0

	for i := 2; i <= N; i++ {
		sum = first + second
		first = second
		second = sum
		fibNums = append(fibNums, second)
	}

	fmt.Println(fibNums)

	return sum
}
