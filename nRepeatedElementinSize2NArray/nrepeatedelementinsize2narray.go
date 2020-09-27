package nRepeatedElementinSize2NArray

func RepeatedNTimesMethod1(A []int) int {

	m := make(map[int]bool)

	for _, v := range A {
		if !m[v] {
			m[v] = true
		} else {
			return v
		}
	}

	return A[0]
}

func RepeatedNTimesMethod2(A []int) int {

	for i := 2; i < len(A); i++ {

		if A[i] == A[i-1] || A[i] == A[i-2] {
			return A[i]
		}

	}

	return A[0]
}
