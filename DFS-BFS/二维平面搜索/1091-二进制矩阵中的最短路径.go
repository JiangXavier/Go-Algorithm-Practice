func shortestPathBinaryMatrix(grid [][]int) int {
	if grid[0][0] == 1 {
		return -1
	}
	n := len(grid)
	dist := make([][]int, n)
	for i, _ := range dist {
		dist[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dist[i][j] = 0x3f3f3f3f
		}
	}
	queue := make([][]int, 0)
	queue = append(queue, []int{0, 0})
	dist[0][0] = 1
	for len(queue) > 0 {
		i, j := queue[0][0], queue[0][1]
		if i == n-1 && j == n-1 {
			return dist[i][j]
		}
		queue = queue[1:]
		direction := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {-1, -1}, {1, 1}, {1, -1}, {-1, 1}}
		for _, d := range direction {
			new_i, new_j := i+d[0], j+d[1]
			if !(0 <= new_i && new_i < n && 0 <= new_j && new_j < n) {
				continue
			}
			if grid[new_i][new_j] == 1 || dist[i][j]+1 >= dist[new_i][new_j] {
				continue
			}
			dist[new_i][new_j] = dist[i][j] + 1
			queue = append(queue, []int{new_i, new_j})
		}
	}
	return -1
}