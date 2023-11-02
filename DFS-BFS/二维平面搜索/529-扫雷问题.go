func updateBoard(board [][]byte, click []int) [][]byte {
    i , j := click[0] , click[1]
    m , n := len(board) , len(board[0])
    v := make([][]bool , m)
    for i , _ := range v{
        v[i] = make([]bool , n)
    }
    if board[i][j] == 'M'{
        board[i][j] = 'X'
    } else{
        dfs(board , i , j , v)
    }
    return board
}
func dfs(board [][]byte , i int , j int , v [][]bool){
    if v[i][j]{return }
    m , n := len(board) , len(board[0])
    cnt := 0
    v[i][j] = true
    direction := [][]int{{1,0},{-1,0},{0,1},{0,-1},{-1,-1},{1,1},{1,-1},{-1,1}}
    for _ , d := range direction{
        new_i , new_j := i + d[0] , j + d[1]
        if 0 <= new_i && new_i < m && 0 <= new_j && new_j < n{
            if board[new_i][new_j] == 'M'{
                cnt++
            }
        }
    }
    if cnt > 0{
        board[i][j] = byte(cnt + '0')
    } else{
        board[i][j] = 'B'
        for _ , d := range direction{
            new_i , new_j := i + d[0] , j + d[1]
            if 0 <= new_i && new_i < m && 0 <= new_j && new_j < n{
                dfs(board , new_i , new_j , v)
            }
        }
    }
}