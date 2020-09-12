package validPalindrome_test

import (
	"GoCode/Iron12/validPalindrome"
	"reflect"
	"testing"
)

type testData struct {
	s   string
	ans bool
}

func generateTestData() []testData {
	data := []testData{
		testData{s: "A man, a plan, a canal: Panama", ans: true},
		testData{s: "race a car", ans: false},
	}
	return data
}

func TestIsPalindromeMethod1(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		actual := validPalindrome.IsPalindromeMethod1(v.s)
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

		actual := validPalindrome.IsPalindromeMethod2(v.s)
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
