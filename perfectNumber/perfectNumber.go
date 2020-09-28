package perfectNumber

func CheckPerfectNumberMethod1(num int) bool {

	if num <= 1 {
		return false
	}

	sum := 1
	tmp := 0

	for i := 2; i*i < num; i++ {
		if num%i == 0 {

			sum += i
			tmp = num / i
			if tmp != i {
				sum += tmp
			}

			if sum > num {
				return false
			}
		}

	}

	if sum != num {
		return false
	}

	return true

}

func CheckPerfectNumberMethod2(num int) bool {

	switch num {
	case 6, 28, 496, 8128, 33550336:
		return true
	default:
		return false
	}

	return false

}
