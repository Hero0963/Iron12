package islandPerimeter

func IslandPerimeterMethod1(grid [][]int) int {
	Ans := 0
	lx := len(grid)
	ly := len(grid[0])

	for i := 0; i < lx; i++ {
		for j := 0; j < ly; j++ {

			// in water region
			if grid[i][j] == 0 {
				continue
			}

			//up
			if j-1 < 0 || grid[i][j-1] == 0 {
				Ans++
			}

			//down
			if j+1 >= ly || grid[i][j+1] == 0 {
				Ans++
			}

			//left
			if i-1 < 0 || grid[i-1][j] == 0 {
				Ans++
			}

			//right
			if i+1 >= lx || grid[i+1][j] == 0 {
				Ans++
			}

		}
	}

	return Ans
}
