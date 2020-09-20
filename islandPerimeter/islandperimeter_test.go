package islandPerimeter_test

import (
	"GoCode/Iron12/islandPerimeter"
	"reflect"
	"testing"
)

type testData struct {
	grid [][]int
	ans  int
}

func generateTestData() []testData {

	grid1 := [][]int{}
	row11 := []int{0, 1, 0, 0}
	row12 := []int{1, 1, 1, 0}
	row13 := []int{0, 1, 0, 0}
	row14 := []int{1, 1, 0, 0}
	grid1 = append(grid1, row11, row12, row13, row14)

	grid2 := [][]int{}
	row21 := []int{1}
	grid2 = append(grid2, row21)

	grid3 := [][]int{}
	row31 := []int{1, 0}
	grid3 = append(grid3, row31)

	data := []testData{
		testData{grid: grid1, ans: 16},
		testData{grid: grid2, ans: 4},
		testData{grid: grid3, ans: 4},
	}
	return data
}

func TestIslandPerimeterMethod1(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		actual := islandPerimeter.IslandPerimeterMethod1(v.grid)
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
