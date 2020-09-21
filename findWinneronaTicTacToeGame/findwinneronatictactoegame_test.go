package findWinneronaTicTacToeGame_test

import (
	"GoCode/Iron12/findWinneronaTicTacToeGame"
	"reflect"
	"testing"
)

type testData struct {
	moves [][]int
	ans   string
}

func generateTestData() []testData {

	moves1 := [][]int{}
	row11 := []int{0, 0}
	row12 := []int{2, 0}
	row13 := []int{1, 1}
	row14 := []int{2, 1}
	row15 := []int{2, 2}
	moves1 = append(moves1, row11, row12, row13, row14, row15)

	moves2 := [][]int{}
	row21 := []int{0, 0}
	row22 := []int{1, 1}
	row23 := []int{0, 1}
	row24 := []int{0, 2}
	row25 := []int{1, 0}
	row26 := []int{2, 0}
	moves2 = append(moves2, row21, row22, row23, row24, row25, row26)

	moves3 := [][]int{}
	row31 := []int{0, 0}
	row32 := []int{1, 1}
	row33 := []int{2, 0}
	row34 := []int{1, 0}
	row35 := []int{1, 2}
	row36 := []int{2, 1}
	row37 := []int{0, 1}
	row38 := []int{0, 2}
	row39 := []int{2, 2}
	moves3 = append(moves3, row31, row32, row33, row34, row35, row36, row37, row38, row39)

	moves4 := [][]int{}
	row41 := []int{0, 0}
	row42 := []int{1, 1}

	moves4 = append(moves4, row41, row42)

	data := []testData{
		testData{moves: moves1, ans: "A"},
		testData{moves: moves2, ans: "B"},
		testData{moves: moves3, ans: "Draw"},
		testData{moves: moves4, ans: "Pending"},
	}
	return data
}

func TestTictactoeMethod1(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		actual := findWinneronaTicTacToeGame.TictactoeMethod1(v.moves)
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

func TestTictactoeMethod2(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		actual := findWinneronaTicTacToeGame.TictactoeMethod2(v.moves)
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

func TestTictactoeMethod3(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		actual := findWinneronaTicTacToeGame.TictactoeMethod3(v.moves)
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
