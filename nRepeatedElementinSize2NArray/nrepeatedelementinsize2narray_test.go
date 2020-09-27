package nRepeatedElementinSize2NArray_test

import (
	"GoCode/Iron12/nRepeatedElementinSize2NArray"
	"reflect"
	"testing"
)

type testData struct {
	A   []int
	ans int
}

func generateTestData() []testData {
	data := []testData{
		testData{A: []int{1, 2, 3, 3}, ans: 3},
		testData{A: []int{2, 1, 2, 5, 3, 2}, ans: 2},
		testData{A: []int{5, 1, 5, 2, 5, 3, 5, 4}, ans: 5},
	}
	return data
}

func TestRepeatedNTimesMethod1(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		actual := nRepeatedElementinSize2NArray.RepeatedNTimesMethod1(v.A)
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

func TestRepeatedNTimesMethod2(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		actual := nRepeatedElementinSize2NArray.RepeatedNTimesMethod2(v.A)
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
