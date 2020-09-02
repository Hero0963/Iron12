package reverseInteger_test

import (
	"GoCode/Iron12/reverseInteger"
	"reflect"
	"testing"
)

type testData struct {
	x   int
	ans int
}

func generateTestData() []testData {
	data := []testData{
		testData{x: 123, ans: 321},
		testData{x: -123, ans: -321},
		testData{x: 120, ans: 21},
	}
	return data
}

func TestReverseMethod1(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		actual := reverseInteger.ReverseMethod1(v.x)
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

func TestReverseMethod2(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		actual := reverseInteger.ReverseMethod2(v.x)
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

func TestReverseMethod2Fix(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		actual := reverseInteger.ReverseMethod2(v.x)
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
