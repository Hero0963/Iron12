package missingNumber_test

import (
	"GoCode/Iron12/missingNumber"
	"reflect"
	"testing"
)

type testData struct {
	nums []int
	ans  int
}

func generateTestData() []testData {
	data := []testData{
		testData{nums: []int{3, 0, 1}, ans: 2},
		testData{nums: []int{9, 6, 4, 2, 3, 5, 7, 0, 1}, ans: 8},
	}
	return data
}

func TestSearchMethod1(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		actual := missingNumber.MissingNumberMethod1(v.nums)
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

func TestSearchMethod2(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		actual := missingNumber.MissingNumberMethod2(v.nums)
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
func TestSearchMethod3(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		actual := missingNumber.MissingNumberMethod3(v.nums)
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
func TestSearchMethod4(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		actual := missingNumber.MissingNumberMethod4(v.nums)
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
