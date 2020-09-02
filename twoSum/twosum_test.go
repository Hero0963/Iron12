package twoSum_test

import (
	"GoCode/Iron12/twoSum"
	"reflect"
	"sort"
	"testing"
)

type testData struct {
	nums   []int
	target int
	ans    []int
}

func generateTestData() []testData {
	data := []testData{
		testData{nums: []int{2, 7, 11, 15}, target: 9, ans: []int{0, 1}},
		testData{nums: []int{3, 2, 4}, target: 6, ans: []int{1, 2}},
		testData{nums: []int{3, 3}, target: 6, ans: []int{0, 1}},
	}
	return data
}

func TestTwoSumMethod1(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		actual := twoSum.TwoSumMethod1(v.nums, v.target)
		expected := v.ans

		sort.Ints(actual)
		sort.Ints(expected)

		if reflect.DeepEqual(expected, actual) {
			t.Log(true)
		} else {
			t.Error("not equal!!")
			t.Error("actual= ", actual)
			t.Error("expected= ", expected)
		}
	}
}

func TestTwoSumMethod2(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		actual := twoSum.TwoSumMethod2(v.nums, v.target)
		expected := v.ans

		sort.Ints(actual)
		sort.Ints(expected)

		if reflect.DeepEqual(expected, actual) {
			t.Log(true)
		} else {
			t.Error("not equal!!")
			t.Error("actual= ", actual)
			t.Error("expected= ", expected)
		}
	}
}
