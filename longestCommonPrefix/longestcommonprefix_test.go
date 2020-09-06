package longestCommonPrefix_test

import (
	"GoCode/Iron12/longestCommonPrefix"
	"reflect"
	"testing"
)

type testData struct {
	strs []string
	ans  string
}

func generateTestData() []testData {
	data := []testData{
		testData{strs: []string{"flower", "flow", "flight"}, ans: "fl"},
		testData{strs: []string{"dog", "racecar", "car"}, ans: ""},
	}
	return data
}

func TestLongestCommonPrefixMethod1(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		actual := longestCommonPrefix.LongestCommonPrefixMethod1(v.strs)
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

func TestLongestCommonPrefixMethod2(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		actual := longestCommonPrefix.LongestCommonPrefixMethod2(v.strs)
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
