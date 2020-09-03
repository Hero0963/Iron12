package palindromeNumber_test

import (
	"GoCode/Iron12/palindromeNumber"
	"reflect"
	"testing"
)

type testData struct {
	x   int
	ans bool
}

func generateTestData() []testData {
	data := []testData{
		testData{x: 121, ans: true},
		testData{x: -121, ans: false},
		testData{x: 10, ans: false},
	}
	return data
}

func TestIsPalindromeMethod1(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		actual := palindromeNumber.IsPalindromeMethod1(v.x)
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

func TestIsPalindromeMethod2(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		actual := palindromeNumber.IsPalindromeMethod2(v.x)
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

func TestIsPalindromeMethod3(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		actual := palindromeNumber.IsPalindromeMethod3(v.x)
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
