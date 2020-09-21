package findWinneronaTicTacToeGame

const actionA string = "X"
const actionB string = "O"

func TictactoeMethod1(moves [][]int) string {

	if len(moves) < 4 {
		return "Pending"
	}

	r := [3][3]string{}

	for i := 0; i < len(moves); i++ {
		x := moves[i][0]
		y := moves[i][1]

		if i%2 == 0 {
			r[x][y] = actionA
		} else {
			r[x][y] = actionB
		}

	}

	if isAWin(r) {
		return "A"
	}

	if isBWin(r) {
		return "B"
	}

	if len(moves) < 9 {
		return "Pending"
	}

	return "Draw"

}

/*
Method2 is learned from
https://leetcode-cn.com/problems/find-winner-on-a-tic-tac-toe-game/solution/java-wei-yun-suan-xiang-jie-shi-yong-wei-yun-suan-/
*/
func TictactoeMethod2(moves [][]int) string {

	a := 0
	b := 0
	l := len(moves)

	ac := []int{7, 56, 448, 73, 146, 292, 273, 84}

	for i := 0; i < l; i++ {
		if (i & 1) == 1 {
			b ^= 1 << (3*moves[i][0] + moves[i][1])
		} else {
			a ^= 1 << (3*moves[i][0] + moves[i][1])
		}

	}

	for _, v := range ac {

		if (a & v) == v {
			return "A"
		}

		if (b & v) == v {
			return "B"
		}
	}

	if l == 9 {
		return "Draw"
	}

	return "Pending"
}

/*
Method 3 id learned from
https://leetcode-cn.com/problems/find-winner-on-a-tic-tac-toe-game/solution/java-shi-jian-zhan-sheng-100-mu-ce-bi-mu-qian-qi-t/
*/
func TictactoeMethod3(moves [][]int) string {

	l := len(moves)
	var count [8]int

	for i := l - 1; i >= 0; i -= 2 {

		// row
		count[moves[i][0]]++

		//row
		count[moves[i][1]+3]++

		//diag1
		if moves[i][0] == moves[i][1] {
			count[6]++
		}

		//diag2
		if moves[i][0]+moves[i][1] == 2 {
			count[7]++
		}

		if count[moves[i][0]] == 3 || count[moves[i][1]+3] == 3 || count[6] == 3 || count[7] == 3 {

			if l%2 == 0 {
				return "B"
			} else {
				return "A"
			}

		}
	}

	if l < 9 {
		return "Pending"
	}

	return "Draw"

}

func isAWin(r [3][3]string) bool {

	if checkCol(r, actionA) {
		return true
	}

	if checkRow(r, actionA) {
		return true
	}

	if checkDiag(r, actionA) {
		return true
	}

	return false
}

func isBWin(r [3][3]string) bool {

	if checkCol(r, actionB) {
		return true
	}

	if checkRow(r, actionB) {
		return true
	}

	if checkDiag(r, actionB) {
		return true
	}

	return false
}

func checkCol(r [3][3]string, s string) bool {

	for i := 0; i < 3; i++ {
		if r[i][0] == s && r[i][1] == s && r[i][2] == s {
			return true
		}
	}

	return false
}

func checkRow(r [3][3]string, s string) bool {

	for j := 0; j < 3; j++ {
		if r[0][j] == s && r[1][j] == s && r[2][j] == s {
			return true
		}
	}

	return false
}

func checkDiag(r [3][3]string, s string) bool {

	if r[0][0] == s && r[1][1] == s && r[2][2] == s {
		return true
	}

	if r[0][2] == s && r[1][1] == s && r[2][0] == s {
		return true
	}

	return false
}
