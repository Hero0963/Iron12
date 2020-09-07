package binarySearch_test

import (
	"GoCode/Iron12/binarySearch"
	"reflect"
	"testing"
)

type testData struct {
	nums   []int
	target int
	ans    int
}

func generateTestData() []testData {
	data := []testData{
		testData{nums: []int{-1, 0, 3, 5, 9, 12}, target: 9, ans: 4},
		testData{nums: []int{-1, 0, 3, 5, 9, 12}, target: 2, ans: -1},
	}
	return data
}

func TestSearchMethod1(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		actual := binarySearch.SearchMethod1(v.nums, v.target)
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

		actual := binarySearch.SearchMethod2(v.nums, v.target)
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
