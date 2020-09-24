package pascalsTriangleII_test

import (
	"GoCode/Iron12/pascalsTriangleII"
	"reflect"
	"testing"
)

type testData struct {
	rowIndex int
	ans      []int
}

func generateTestData() []testData {
	data := []testData{
		testData{rowIndex: 3, ans: []int{1, 3, 3, 1}},
		testData{rowIndex: 0, ans: []int{1}},
		testData{rowIndex: 1, ans: []int{1, 1}},
	}
	return data
}

func TestGetRowMethod1(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		actual := pascalsTriangleII.GetRowMethod1(v.rowIndex)
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
