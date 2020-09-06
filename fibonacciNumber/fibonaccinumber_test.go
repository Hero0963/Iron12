package fibonacciNumber_test

import (
	"GoCode/Iron12/fibonacciNumber"
	"reflect"
	"testing"
)

type testData struct {
	N   int
	ans int
}

func generateTestData() []testData {
	data := []testData{
		testData{N: 2, ans: 1},
		testData{N: 3, ans: 2},
		testData{N: 4, ans: 3},
		testData{N: 5, ans: 5},
	}
	return data
}

func TestFibMethod1(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		actual := fibonacciNumber.FibMethod1(v.N)
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

func TestFibMethod2(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		actual := fibonacciNumber.FibMethod2(v.N)
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

func TestFibMethod3(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		actual := fibonacciNumber.FibMethod3(v.N)
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

func TestFibMethod4(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		actual := fibonacciNumber.FibMethod4(v.N)
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
