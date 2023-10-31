//DFS
//分别逆向搜索能从什么地方流出，然后求出交集
func pacificAtlantic(heights [][]int) [][]int {
	ans := make([][]int, 0)
	direction := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	m, n := len(heights), len(heights[0])
	v := make([][]bool, m)
	tai := make([][]bool, m)
	da := make([][]bool, m)
	for i, _ := range v {
		v[i] = make([]bool, n)
		tai[i] = make([]bool, n)
		da[i] = make([]bool, n)
	}
	var dfs func(i, j int, shu [][]bool)
	dfs = func(i, j int, shu [][]bool) {
		if v[i][j] {
			return
		}
		v[i][j] = true
		shu[i][j] = true
		for _, d := range direction {
			new_i, new_j := i+d[0], j+d[1]
			if 0 <= new_i && new_i < m && 0 <= new_j && new_j < n && heights[new_i][new_j] >= heights[i][j] {
				dfs(new_i, new_j, shu)
			}
		}
	}
	for i := 0; i < m; i++ {
		dfs(i, 0, tai)
	}
	for i := 0; i < n; i++ {
		dfs(0, i, tai)
	}
	for i, _ := range v {
		v[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		dfs(i, n-1, da)
	}
	for i := 0; i < n; i++ {
		dfs(m-1, i, da)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if tai[i][j] && da[i][j] {
				ans = append(ans, []int{i, j})
			}
		}
	}
	return ans
}

// BFS
func pacificAtlantic(heights [][]int) [][]int {
	ans := make([][]int, 0)
	direction := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	m, n := len(heights), len(heights[0])
	v := make([][]bool, m)
	tai := make([][]bool, m)
	da := make([][]bool, m)
	for i, _ := range v {
		v[i] = make([]bool, n)
		tai[i] = make([]bool, n)
		da[i] = make([]bool, n)
	}
	var dfs func(i, j int, shu [][]bool)
	dfs = func(i, j int, shu [][]bool) {
		if v[i][j] {
			return
		}
		queue := make([][]int, 0)
		queue = append(queue, []int{i, j})
		for len(queue) > 0 {
			q := queue[0]
			queue = queue[1:]
			if v[q[0]][q[1]] {
				continue
			}
			v[q[0]][q[1]] = true
			shu[q[0]][q[1]] = true
			for _, d := range direction {
				new_i, new_j := q[0]+d[0], q[1]+d[1]
				if 0 <= new_i && new_i < m && 0 <= new_j && new_j < n && heights[new_i][new_j] >= heights[q[0]][q[1]] {
					queue = append(queue, []int{new_i, new_j})
				}
			}
		}
	}
	for i := 0; i < m; i++ {
		dfs(i, 0, tai)
	}
	for i := 0; i < n; i++ {
		dfs(0, i, tai)
	}
	for i, _ := range v {
		v[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		dfs(i, n-1, da)
	}
	for i := 0; i < n; i++ {
		dfs(m-1, i, da)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if tai[i][j] && da[i][j] {
				ans = append(ans, []int{i, j})
			}
		}
	}
	return ans
}