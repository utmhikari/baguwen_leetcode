package solution_151_200

//200. 岛屿数量
//给你一个由'1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
//
//岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
//
//此外，你可以假设该网格的四条边均被水包围。



var cnt200 int


func dfsLand200(x int, y int, h int, w int, grid *[][]byte, visited *[][]bool) {
	if x < 0 || x >= h || y < 0 || y >= w || (*visited)[x][y] {
		return
	}
	(*visited)[x][y] = true

	if string((*grid)[x][y]) == "0" {
		return
	}

	dfsLand200(x + 1, y, h, w, grid, visited)
	dfsLand200(x - 1, y, h, w, grid, visited)
	dfsLand200(x, y + 1, h, w, grid, visited)
	dfsLand200(x, y - 1, h, w, grid, visited)
}


func numIslands200(grid [][]byte) int {
	h := len(grid)
	if h == 0 {
		return 0
	}

	w := len(grid[0])
	if w == 0 {
		return 0
	}

	visited := make([][]bool, h)
	for i := 0; i < h; i++ {
		visited[i] = make([]bool, w)
	}

	cnt200 = 0

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if visited[i][j] || string(grid[i][j]) == "0" {
				continue
			}
			cnt200++
			dfsLand200(i, j, h, w, &grid, &visited)
		}
	}

	return cnt200
}



