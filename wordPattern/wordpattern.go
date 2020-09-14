package wordPattern

import "strings"

func WordPatternMethod1(pattern string, str string) bool {

	ss := strings.Split(str, " ")
	mp := make(map[byte]string)
	ms := make(map[string]byte)

	var p byte
	var s string
	var okp, oks bool

	if len(pattern) != len(ss) {
		return false
	}

	for i := 0; i < len(pattern); i++ {
		p, s = pattern[i], ss[i]
		_, okp = mp[p]
		_, oks = ms[s]

		if !okp && !oks {
			mp[p] = s
			ms[s] = p
		}

		if mp[p] != s || ms[s] != p {
			return false
		}

	}

	return true
}
