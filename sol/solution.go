package sol

type Pair struct {
	row, col int
}

func maxAreaOfIsland(grid [][]int) int {
	ROW := len(grid)
	COL := len(grid[0])
	visit := make(map[Pair]struct{})
	var dfs func(row int, col int) int
	dfs = func(row int, col int) int {
		if row < 0 || row >= ROW || col < 0 || col >= COL ||
			grid[row][col] == 0 {
			return 0
		}
		if _, ok := visit[Pair{row: row, col: col}]; ok {
			return 0
		}
		visit[Pair{row: row, col: col}] = struct{}{}
		return 1 + dfs(row-1, col) + dfs(row+1, col) + dfs(row, col-1) + dfs(row, col+1)
	}
	area := 0
	for row := 0; row < ROW; row++ {
		for col := 0; col < COL; col++ {
			total := dfs(row, col)
			if total > area {
				area = total
			}
		}
	}
	return area
}
