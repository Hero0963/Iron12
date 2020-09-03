package palindromeNumber

func IsPalindromeMethod1(x int) bool {
	if x < 0 {
		return false
	}

	if x == 0 {
		return true
	}

	var nums []int
	var r int

	for x > 0 {
		r = x % 10
		nums = append(nums, r)
		x = x / 10
	}

	head := 0
	tail := len(nums) - 1

	for head < tail {
		if nums[head] != nums[tail] {
			return false
		}
		head++
		tail--
	}
	return true
}

func IsPalindromeMethod2(x int) bool {
	if x < 0 {
		return false
	}

	if x == 0 {
		return true
	}

	ori := x
	rev := 0
	r := 0

	for x > 0 {
		r = x % 10
		rev = rev*10 + r
		x /= 10

	}

	if ori == rev {
		return true
	}

	return false
}

func IsPalindromeMethod3(x int) bool {
	if x < 0 {
		return false
	}

	if x == 0 {
		return true
	}

	if x%10 == 0 {
		return false
	}

	rev := 0
	r := 0

	for {
		r = x % 10
		rev = rev*10 + r
		x /= 10

		if rev >= x {
			break
		}

	}

	return x == rev || x == rev/10
}
