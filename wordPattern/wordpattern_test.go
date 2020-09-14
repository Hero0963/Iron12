package wordPattern_test

import (
	"GoCode/Iron12/wordPattern"
	"reflect"
	"testing"
)

type testData struct {
	pattern string
	str     string
	ans     bool
}

func generateTestData() []testData {
	data := []testData{
		testData{pattern: "abba", str: "dog cat cat dog", ans: true},
		testData{pattern: "abba", str: "dog cat cat fish", ans: false},
		testData{pattern: "aaaa", str: "dog cat cat dog", ans: false},
	}
	return data
}

func TestWordPatternMethod1(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		actual := wordPattern.WordPatternMethod1(v.pattern, v.str)
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
