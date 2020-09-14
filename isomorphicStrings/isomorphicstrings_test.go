package isomorphicStrings_test

import (
	"GoCode/Iron12/isomorphicStrings"
	"reflect"
	"testing"
)

type testData struct {
	s   string
	t   string
	ans bool
}

func generateTestData() []testData {
	data := []testData{
		testData{s: "egg", t: "add", ans: true},
		testData{s: "foo", t: "bar", ans: false},
		testData{s: "paper", t: "title", ans: true},
	}
	return data
}

func TestIsIsomorphicMethod1(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		actual := isomorphicStrings.IsIsomorphicMethod1(v.s, v.t)
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

func TestIsIsomorphicMethod2(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		actual := isomorphicStrings.IsIsomorphicMethod2(v.s, v.t)
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
