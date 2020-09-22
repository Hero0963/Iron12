package climbingStairs_test

import (
	"GoCode/Iron12/climbingStairs"
	"reflect"
	"testing"
)

type testData struct {
	n   int
	ans int
}

func generateTestData() []testData {
	data := []testData{
		testData{n: 2, ans: 2},
		testData{n: 3, ans: 3},
	}
	return data
}

func TestClimbStairsMethod1(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		actual := climbingStairs.ClimbStairsMethod1(v.n)
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
