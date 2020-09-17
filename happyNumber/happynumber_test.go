package happyNumber_test

import (
	"GoCode/Iron12/happyNumber"
	"reflect"
	"testing"
)

type testData struct {
	n   int
	ans bool
}

func generateTestData() []testData {
	data := []testData{
		testData{n: 19, ans: true},
		testData{n: 4, ans: false},
		testData{n: 322, ans: false},
	}
	return data
}

func TestIsHappyMethod1(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		actual := happyNumber.IsHappyMethod1(v.n)
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

func TestIsHappyMethod2(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		actual := happyNumber.IsHappyMethod2(v.n)
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

func TestIsHappyMethod3(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		actual := happyNumber.IsHappyMethod3(v.n)
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
