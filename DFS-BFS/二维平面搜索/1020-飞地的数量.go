//DFS
func numEnclaves(board [][]int) int {
    m, n := len(board), len(board[0])
    ans := 0
	if m < 3 || n < 3 {
		return 0
	}
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if !(0 <= i && i < m && 0 <= j && j < n) || !(board[i][j] == 1) {
			return
		}
		board[i][j] = 2
		dfs(i, j+1)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i-1, j)
	}
	for i := 0; i < m; i++ {
		dfs(i, 0)
		dfs(i, n-1)
	}
	for j := 0; j < n; j++ {
		dfs(0, j)
		dfs(m-1, j)
	}
	for i, row := range board {
		for j := range row {
			if board[i][j] == 1 {
				ans ++
			} 
		}
	}
    return ans
}

//BFS
func numEnclaves(board [][]int) int {
	m , n := len(board) , len(board[0])
	ans :=  0
	if m < 3 || n < 3 {return 0}
	v := make([][]bool , m)
	for i , _ := range v{
		v[i] = make([]bool , n)
	}
	for i := 0 ; i < m ; i++{
		bfs(i,0,board,v)
		bfs(i,n-1,board,v)
	}
	for j := 0; j < n ;j++{
		bfs(0,j,board,v)
		bfs(m-1,j,board,v)
	}
	for i , row := range board{
		for j := range row{
			if board[i][j] == 1{
				ans ++
			}
		}
	}
	return ans
}
func bfs(i,j int , board [][]int, v [][]bool){
	m , n := len(board) , len(board[0])
	direction := [][]int{{1,0},{-1,0},{0,1},{0,-1}}
	if v[i][j] {return }
	if !(0 <= i && i < m && 0 <= j && j < n) || !(board[i][j] == 1) {
		return
	}
	queue := make([][]int , 0)
	queue = append(queue , []int{i,j})
	for len(queue) > 0{
		q := queue[0]
		queue = queue[1:]
		board[q[0]][q[1]] = 2
		//关键剪枝
		if v[q[0]][q[1]]{
			continue
		}
		v[q[0]][q[1]] = true
		for _ , new_d := range direction{
			new_i , new_j := q[0] + new_d[0], q[1] + new_d[1]
			if 0 <= new_i && new_i < m && 0 <= new_j && new_j < n && board[new_i][new_j] == 1{
				queue = append(queue , []int{new_i , new_j})
			}
		}
	}
}