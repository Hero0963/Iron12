package reverseInteger

import (
	"fmt"
	"math"
	"strconv"
)

func ReverseMethod1(x int) int {

	sign := 1
	if x < 0 {
		sign = -1
		x = -x
	}

	s := strconv.Itoa(x)
	bs := []byte(s)

	head := 0
	tail := len(bs) - 1

	for head < tail {
		bs[head], bs[tail] = bs[tail], bs[head]
		head++
		tail--
	}

	rs := string(bs)

	if sign == -1 {
		rs = "-" + rs
	}

	y, err := strconv.Atoi(rs)
	if err != nil {
		fmt.Println(err, rs)
	}

	if (y > math.MaxInt32) || (y < math.MinInt32) {
		return 0
	}

	return y
}

func ReverseMethod2(x int) int {

	y := 0
	r := 0
	sign := 1

	if x < 0 {
		sign = -1
		x = -x
	}

	for x > 0 {
		r = x % 10
		y = y*10 + r
		x = x / 10
	}

	y *= sign

	if (y > math.MaxInt32) || (y < math.MinInt32) {
		return 0
	}

	return y
}

func ReverseMethod2Fix(x int) int {

	prey := 0
	y := 0
	r := 0
	sign := 1

	if x < 0 {
		sign = -1
		x = -x
	}

	for x > 0 {
		r = x % 10
		y = y*10 + r

		if ((y - r) / 10) != prey {
			return 0
		}

		prey = y
		x = x / 10
	}

	y *= sign

	if (y > math.MaxInt32) || (y < math.MinInt32) {
		return 0
	}

	return y
}
