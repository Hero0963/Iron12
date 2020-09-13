package plusOne

func PlusOneMethod1(digits []int) []int {

	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		} else {
			digits[i] = 0
		}
	}

	digits = append(digits, 0)
	digits[0] = 1

	return digits
}

func PlusOneMethod2(digits []int) []int {

	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		} else {
			digits[i] = 0
		}
	}

	digits = append(digits, 0)
	digits[0] = 1

	return digits
}
