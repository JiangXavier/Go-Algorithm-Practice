//DFS
func maxAreaOfIsland(grid [][]int) int {
	direction := [][]int{{1, 0}, {0, -1}, {-1, 0}, {0, 1}}
	m, n := len(grid), len(grid[0])
	v := make([][]bool, m)
	ans := 0
	for i := range v {
		v[i] = make([]bool, n)
	}
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		v[i][j] = true
		a := 1
		for _, new_d := range direction {
			new_i, new_j := i+new_d[0], j+new_d[1]
			if 0 <= new_i && new_i < m && 0 <= new_j && new_j < n && !v[new_i][new_j] && grid[new_i][new_j] == 1 {
				a += dfs(new_i, new_j)
			}
		}
		return a
	}

	for i, row := range grid {
		for j := range row {
			if grid[i][j] == 1 && !v[i][j] {
				ans = max(ans, dfs(i, j))
			}
		}
	}
	return ans
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
// BFS
func maxAreaOfIsland(grid [][]int) int {
    direction := [][]int{{1, 0}, {0, -1}, {-1, 0}, {0, 1}}
	m, n := len(grid), len(grid[0])
	v := make([][]bool, m)
    ans := 0
	for i := range v {
		v[i] = make([]bool, n)
	}
    queue := make([][]int , 0) 

    for i , row := range grid{
        for j := range row{
            if grid[i][j] == 1 && !v[i][j]{
                tmp := 0
                queue = append(queue , []int{i,j})
                v[i][j] = true
                for len(queue) > 0{
                    q := queue[0]
                    queue = queue[1:]
                    tmp += 1
                    for _ , new_d := range direction{
                        new_i , new_j := q[0] + new_d[0], q[1] + new_d[1]
                        if 0 <= new_i && new_i < m && 0 <= new_j && new_j < n && !v[new_i][new_j] && grid[new_i][new_j] == 1{
                            queue = append(queue , []int{new_i , new_j})
                            v[new_i][new_j] = true
                        }
                    }
                }
                ans = max(ans , tmp)
            }
        }
    }
    return ans
}
func max(a , b int) int {
    if a > b {return a}
    return b
}