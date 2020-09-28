package perfectNumber_test

import (
	"GoCode/Iron12/perfectNumber"
	"reflect"
	"testing"
)

type testData struct {
	num int
	ans bool
}

func generateTestData() []testData {
	data := []testData{
		testData{num: 28, ans: true},
	}
	return data
}

func TestCheckPerfectNumberMethod1(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		actual := perfectNumber.CheckPerfectNumberMethod1(v.num)
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

func TestCheckPerfectNumberMethod2(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		actual := perfectNumber.CheckPerfectNumberMethod2(v.num)
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
