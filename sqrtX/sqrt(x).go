package sqrtX

import "math"

func MySqrtMethod1(x int) int {

	i := 0

	for i*i <= x {
		i++
	}

	return i - 1
}

func MySqrtMethod2(x int) int {

	if x <= 1 {
		return x
	}

	head := 0
	tail := x/2 + 1
	mid := 0

	for head < tail {
		mid = head + (tail-head)/2

		if mid*mid == x {
			return mid
		} else if mid*mid > x {
			tail = mid
		} else {
			head = mid + 1
		}

	}

	return tail - 1
}

func MySqrtMethod3(x int) int {

	if x <= 1 {
		return x
	}

	r := x
	for r*r > x {
		r = (r + x/r) / 2
	}
	return r
}

func MySqrtMethod4(x int) int {
	return int(math.Sqrt(float64(x)))
}
