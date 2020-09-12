package validPalindrome

func IsPalindromeMethod1(s string) bool {
	head := 0
	tail := len(s) - 1

	for head < tail {
		if !isAlphabet(s[head]) && !isNumeric(s[head]) {
			head++
			continue
		}
		if !isAlphabet(s[tail]) && !isNumeric(s[tail]) {
			tail--
			continue
		}
		if convertToLowerCase(s[head]) != convertToLowerCase(s[tail]) {
			return false
		}

		head++
		tail--
	}

	return true

}

func IsPalindromeMethod2(s string) bool {

	bs := []byte{}
	var b byte
	for _, v := range s {
		b = byte(v)
		if isAlphabet(b) || isNumeric(b) {
			bs = append(bs, convertToLowerCase(b))
		}
	}

	head := 0
	tail := len(bs) - 1

	for head < tail {
		if bs[head] != bs[tail] {
			return false
		}

		head++
		tail--
	}

	return true

}

func isAlphabet(b byte) bool {
	if b >= 'a' && b <= 'z' {
		return true
	}

	if b >= 'A' && b <= 'Z' {
		return true
	}

	return false
}

func isNumeric(b byte) bool {
	if b >= '0' && b <= '9' {
		return true
	}
	return false
}

func convertToLowerCase(b byte) byte {

	if b >= 'A' && b <= 'Z' {
		b = b - 'A' + 'a'
	}

	return b

}
