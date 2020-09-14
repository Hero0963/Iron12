package isomorphicStrings

func IsIsomorphicMethod1(s string, t string) bool {

	var arrs, arrt [128]int
	var bs, bt byte

	for i := 0; i < len(s); i++ {
		bs, bt = s[i], t[i]
		if arrs[bs] != arrt[bt] {
			return false
		}
		arrs[bs], arrt[bt] = i+1, i+1
	}

	return true

}

func IsIsomorphicMethod2(s string, t string) bool {

	var arys, aryt [128]byte
	var bs, bt byte

	for i := 0; i < len(s); i++ {
		bs, bt = s[i], t[i]

		if arys[bs] == 0 && aryt[bt] == 0 {
			arys[bs], aryt[bt] = bt, bs
		}

		if arys[bs] != bt || aryt[bt] != bs {
			return false
		}

	}

	return true

}
