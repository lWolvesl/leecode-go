package lee2307

import "sort"

func deleteGreatestValue(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	for i := 0; i < m; i++ {
		sort.Ints(grid[i])
	}
	res := 0
	for j := 0; j < n; j++ {
		mx := 0
		for i := 0; i < m; i++ {
			mx = max(mx, grid[i][j])
		}
		res += mx
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
