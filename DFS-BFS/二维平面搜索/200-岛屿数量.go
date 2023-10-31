//DFS
func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	ans := 0
	v := make([][]bool, m)
	for i, _ := range v {
		v[i] = make([]bool, n)
	}
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if !(0 <= i && i < m && 0 <= j && j < n) || v[i][j] || grid[i][j] == '0' {
			return
		}
		v[i][j] = true
		dfs(i, j+1)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i-1, j)
	}

	for i, row := range grid {
		for j := range row {
			if grid[i][j] == '1' && !v[i][j] {
				ans += 1
				dfs(i, j)
			}
		}
	}
	return ans
}

//BFS
func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	ans := 0
	v := make([][]bool, m)
	for i, _ := range v {
		v[i] = make([]bool, n)
	}

	for i, row := range grid {
		for j := range row {
			if grid[i][j] == '1' && !v[i][j] {
				ans += 1
				queue := make([][]int, 0)
				direction := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
				queue = append(queue, []int{i, j})
				for len(queue) > 0 {
					q := queue[0]
					queue = queue[1:]
					for _, d := range direction {
						new_i, new_j := q[0]+d[0], q[1]+d[1]
						if !(0 <= new_i && new_i < m && 0 <= new_j && new_j < n) || v[new_i][new_j] || grid[new_i][new_j] == '0' {
							continue
						}
						queue = append(queue, []int{new_i, new_j})
						v[new_i][new_j] = true
					}

				}
			}
		}
	}
	return ans
}