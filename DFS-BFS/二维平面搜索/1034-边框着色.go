func colorBorder(grid [][]int, row int, col int, color int) [][]int {
    m , n := len(grid) , len(grid[0])
    direction := [][]int{{1,0},{-1,0},{0,1},{0,-1}}
    c := grid[row][col]
    v := make([][]bool , m)
    for i , _ := range v{
        v[i] = make([]bool , n)
    }
    var dfs func(i,j int)
    dfs = func(i,j int){
        if !(0<=i&&i<m&&0<=j&&j<n){return }
        if v[i][j] {return }
        v[i][j] = true
        if (grid[i][j] != c) && (grid[i][j] != 0) {return }
        aa := 0
        for _ , d := range direction{
            new_i , new_j := i + d[0] , j + d[1]
            if (0<=new_i&&new_i<m&&0<=new_j&&new_j<n) && (grid[new_i][new_j] == c || grid[new_i][new_j] == 0){
                aa++
            }
        }   
        if aa == 4{grid[i][j] = 0}     
        dfs(i,j+1)
        dfs(i,j-1)
        dfs(i+1,j)
        dfs(i-1,j)
    }
    dfs(row,col)
    for i , row := range grid{
        for j := range row{
            if grid[i][j] == c && v[i][j]{
                grid[i][j] = color
            } else if grid[i][j] == 0{
                grid[i][j] = c
            }
        }
    }
    return grid
}