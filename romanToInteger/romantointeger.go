package romanToInt

var m = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func RomanToIntMethod1(s string) int {
	sum := 0
	temp := 0
	last := 0

	for i := len(s) - 1; i >= 0; i-- {
		temp = m[s[i]]
		sign := 1
		if temp < last {
			sign = -1
		}

		sum = sum + (temp * sign)
		last = temp
	}

	return sum

}
