package plusOne_test

import (
	"GoCode/Iron12/plusOne"
	"reflect"
	"testing"
)

type testData struct {
	digits []int
	ans    []int
}

func generateTestData() []testData {
	data := []testData{
		testData{digits: []int{1, 2, 3}, ans: []int{1, 2, 4}},
		testData{digits: []int{4, 3, 2, 1}, ans: []int{4, 3, 2, 2}},
		testData{digits: []int{0}, ans: []int{1}},
	}
	return data
}

func TestPlusOneMethod1(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		actual := plusOne.PlusOneMethod1(v.digits)
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

func TestPlusOneMethod2(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		actual := plusOne.PlusOneMethod2(v.digits)
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
