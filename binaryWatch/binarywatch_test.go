package binaryWatch_test

import (
	"GoCode/Iron12/binaryWatch"
	"reflect"
	"sort"
	"testing"
)

type testData struct {
	num int
	ans []string
}

func generateTestData() []testData {
	data := []testData{
		testData{num: 1, ans: []string{"1:00", "2:00", "4:00", "8:00", "0:01", "0:02", "0:04", "0:08", "0:16", "0:32"}},
	}
	return data
}

func TestReadBinaryWatchMethod2(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		actual := binaryWatch.ReadBinaryWatchMethod2(v.num)
		expected := v.ans

		sort.Strings(actual)
		sort.Strings(expected)

		if reflect.DeepEqual(expected, actual) {
			t.Log(true)
		} else {
			t.Error("not equal!!")
			t.Error("actual= ", actual)
			t.Error("expected= ", expected)
		}
	}
}

func TestReadBinaryWatchMethod1(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		actual := binaryWatch.ReadBinaryWatchMethod1(v.num)
		expected := v.ans

		sort.Strings(actual)
		sort.Strings(expected)

		if reflect.DeepEqual(expected, actual) {
			t.Log(true)
		} else {
			t.Error("not equal!!")
			t.Error("actual= ", actual)
			t.Error("expected= ", expected)
		}
	}
}
