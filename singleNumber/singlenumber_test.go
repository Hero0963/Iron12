package singleNumber_test

import (
	"GoCode/Iron12/singleNumber"
	"reflect"
	"testing"
)

type testData struct {
	nums []int
	ans  int
}

func generateTestData() []testData {
	data := []testData{
		testData{nums: []int{2, 2, 1}, ans: 1},
		testData{nums: []int{4, 1, 2, 1, 2}, ans: 4},
	}
	return data
}

func TestSingleNumberMethod1(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		actual := singleNumber.SingleNumberMethod1(v.nums)
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

func TestSingleNumberMethod2(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		actual := singleNumber.SingleNumberMethod2(v.nums)
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
