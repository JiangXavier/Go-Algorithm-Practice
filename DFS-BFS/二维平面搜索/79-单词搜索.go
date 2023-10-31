//DFS
//defer + 匿名函数可在回溯结束后恢复状态
func exist(board [][]byte, word string) bool {
	direction := [][]int{{1, 0}, {0, -1}, {-1, 0}, {0, 1}}
	m, n := len(board), len(board[0])
	v := make([][]bool, m)
	for i := range v {
		v[i] = make([]bool, n)
	}
	var dfs func(i, j, begin int) bool
	dfs = func(i, j, begin int) bool {
		if board[i][j] != word[begin] {
			return false
		}
		if begin == len(word)-1 {
			return true
		}
		v[i][j] = true
		//defer func() {v[i][j] = false}()
		for _, new_d := range direction {
			new_i, new_j := new_d[0]+i, new_d[1]+j
			if new_i >= 0 && new_i < m && new_j >= 0 && new_j < n && !v[new_i][new_j] {
				if dfs(new_i, new_j, begin+1) {
					return true
				}
			}
		}
		v[i][j] = false
		return false
	}
	for i, row := range board {
		for j := range row {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}