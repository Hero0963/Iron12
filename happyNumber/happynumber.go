package happyNumber

func IsHappyMethod1(n int) bool {
	m := make(map[int]bool)
	m[n] = true

	for n != 1 {
		n = helper(n)
		if m[n] == true {
			return false
		}

		m[n] = true
	}

	return true
}

func IsHappyMethod2(n int) bool {
	slow := n
	fast := helper(n)

	for fast != 1 && fast != slow {
		slow = helper(slow)
		fast = helper(helper(fast))
	}

	return fast == 1
}

func IsHappyMethod3(n int) bool {

	for n != 1 {
		n = helper(n)
		if n < 10 {
			if n == 1 || n == 7 {
				return true
			}
			return false
		}
	}

	return true
}

func helper(n int) (next int) {

	r := 0
	for n != 0 {
		r = n % 10
		next += r * r
		n = n / 10
	}
	return next
}
