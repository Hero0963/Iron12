package longestCommonPrefix

func LongestCommonPrefixMethod1(strs []string) string {

	if len(strs) == 0 {
		return ""
	}

	s := ""

	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if (i >= len(strs[j])) || (strs[0][i] != strs[j][i]) {
				return s
			}
		}
		s += string(strs[0][i])
	}

	return s
}

func LongestCommonPrefixMethod2(strs []string) string {

	if len(strs) == 0 {
		return ""
	}

	s := ""

	minL := len(strs[0])

	for _, v := range strs {

		if len(v) < minL {
			minL = len(v)
		}

	}

	for i := 0; i < minL; i++ {
		for j := 1; j < len(strs); j++ {
			if strs[0][i] != strs[j][i] {
				return s
			}
		}
		s += string(strs[0][i])
	}

	return s
}
