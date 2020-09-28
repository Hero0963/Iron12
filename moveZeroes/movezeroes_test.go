package moveZeroes_test

import (
	"GoCode/Iron12/moveZeroes"
	"reflect"
	"testing"
)

type testData struct {
	nums []int
	ans  []int
}

func generateTestData() []testData {
	data := []testData{
		testData{nums: []int{0, 1, 0, 3, 12}, ans: []int{1, 3, 12, 0, 0}},
	}
	return data
}

func TestMajorityElementMethod1(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		moveZeroes.MoveZeroesMethod1(v.nums)
		actual := v.nums

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

		moveZeroes.MoveZeroesMethod2(v.nums)
		actual := v.nums

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

func TestMajorityElementMethod3(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		moveZeroes.MoveZeroesMethod3(v.nums)
		actual := v.nums

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
