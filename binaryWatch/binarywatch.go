package binaryWatch

import "strconv"

func ReadBinaryWatchMethod1(num int) []string {

	var leds [10]int
	var tmp, pos, r int
	var ans []string

	for i := 0; i < 1024; i++ {

		leds = [10]int{}
		tmp, pos, r = i, len(leds)-1, 0

		for tmp > 0 {
			r = tmp % 2

			if r != 0 {
				leds[pos]++
			}
			tmp /= 2
			pos--
		}

		if !isCorrectNum(leds, num) {
			continue
		}

		h, m := convertLEDToTime(leds)

		if isLegalTime(h, m) {

			s := convertToRequiredFormat(h, m)
			ans = append(ans, s)
		}

	}

	return ans

}

func ReadBinaryWatchMethod2(n int) []string {

	var ans []string
	leds := make([]bool, 10)

	var dfs func(int, int)
	dfs = func(idx, n int) {
		var h, m int
		if n == 0 {
			m, h = convertToTimeValue(leds[:6]), convertToTimeValue(leds[6:])

			if isLegalTime(h, m) {
				s := convertToRequiredFormat(h, m)
				ans = append(ans, s)
			}

			return
		}

		for i := idx; i < len(leds)-n+1; i++ {
			leds[i] = true
			dfs(i+1, n-1)
			leds[i] = false
		}
	}

	dfs(0, n)

	return ans
}

var bs = []int{1, 2, 4, 8, 16, 32}

func convertToTimeValue(leds []bool) (sum int) {

	for i, v := range leds {
		if v == true {
			sum += bs[i]
		}
	}

	return sum
}

func isCorrectNum(leds [10]int, num int) bool {

	c := 0

	for _, v := range leds {
		if v != 0 {
			c++
		}
	}

	if c != num {
		return false
	}

	return true

}

func convertLEDToTime(leds [10]int) (hours int, minutes int) {

	tmp := 0

	for i, v := range leds {

		if v == 0 {
			continue
		}

		if i < 6 {
			tmp = 1 << uint(i)
			minutes += tmp
		} else {
			tmp = 1 << uint((i - 6))
			hours += tmp
		}

	}

	return hours, minutes
}

func isLegalTime(hours int, minutes int) bool {

	if hours < 12 && minutes < 60 {
		return true
	}

	return false

}

func convertToRequiredFormat(hours int, minutes int) (s string) {

	ts := ""
	if minutes < 10 {
		ts = "0"
	}

	s = strconv.Itoa(hours) + ":" + ts + strconv.Itoa(minutes)

	return s
}
