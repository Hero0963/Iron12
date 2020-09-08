package sqrtX_test

import (
	"GoCode/Iron12/sqrtX"
	"reflect"
	"testing"
)

type testData struct {
	x   int
	ans int
}

func generateTestData() []testData {
	data := []testData{
		testData{x: 4, ans: 2},
		testData{x: 8, ans: 2},
	}
	return data
}

func TestMySqrtMethod1(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		actual := sqrtX.MySqrtMethod1(v.x)
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

func TestMySqrtMethod2(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		actual := sqrtX.MySqrtMethod2(v.x)
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

func TestMySqrtMethod3(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		actual := sqrtX.MySqrtMethod3(v.x)
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

func TestMySqrtMethod4(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		actual := sqrtX.MySqrtMethod4(v.x)
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
