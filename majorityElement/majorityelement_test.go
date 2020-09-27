package majorityElement_test

import (
	"GoCode/Iron12/majorityElement"
	"reflect"
	"testing"
)

type testData struct {
	nums []int
	ans  int
}

func generateTestData() []testData {
	data := []testData{
		testData{nums: []int{3, 2, 3}, ans: 3},
		testData{nums: []int{2, 2, 1, 1, 1, 2, 2}, ans: 2},
	}
	return data
}

func TestMajorityElementMethod1(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		actual := majorityElement.MajorityElementMethod1(v.nums)
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

func TestMajorityElementMethod2(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		actual := majorityElement.MajorityElementMethod2(v.nums)
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
