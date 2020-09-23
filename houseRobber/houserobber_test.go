package houseRobber_test

import (
	"GoCode/Iron12/houseRobber"
	"reflect"
	"testing"
)

type testData struct {
	nums []int
	ans  int
}

func generateTestData() []testData {
	data := []testData{
		testData{nums: []int{1, 2, 3, 1}, ans: 4},
		testData{nums: []int{2, 7, 9, 3, 1}, ans: 12},
	}
	return data
}

func TestRobMethod1(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		actual := houseRobber.RobMethod1(v.nums)
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
