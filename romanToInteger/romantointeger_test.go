package romanToInt_test

import (
	romanToInt "GoCode/Iron12/romanToInteger"
	"reflect"
	"testing"
)

type testData struct {
	s   string
	ans int
}

func generateTestData() []testData {
	data := []testData{
		testData{s: "III", ans: 3},
		testData{s: "IV", ans: 4},
		testData{s: "IX", ans: 9},
		testData{s: "LVIII", ans: 58},
		testData{s: "MCMXCIV", ans: 1994},
	}
	return data
}

func TestRomanToIntMethod1(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		actual := romanToInt.RomanToIntMethod1(v.s)
		expected := v.ans

		if reflect.DeepEqual(expected, actual) {
			t.Log(true)
		} else {
			t.Error("not equal!!")
			t.Error("actual= ", actual)
			t.Error("expected= ", expected)
		}
	}
}
