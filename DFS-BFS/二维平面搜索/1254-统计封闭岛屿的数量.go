func closedIsland(grid [][]int) int {
    m , n := len(grid) , len(grid[0])
    direction := [][]int{{1,0},{-1,0},{0,1},{0,-1}}
    if m < 3 || n < 3{return 0}
    ans := 0
    v := make([][]bool , m)
    for i , _ := range v{
        v[i] = make([]bool , n)
    }
    var dfs func(i,j int)
    dfs = func(i,j int){
        if !(0 <= i && i < m && 0<=j && j < n){return }
        if grid[i][j] == 1 {return }
        grid[i][j] = 1
        dfs(i,j+1)
        dfs(i,j-1)
        dfs(i+1,j)
        dfs(i-1,j)
    }
    for i := 0;i<m;i++{
        dfs(i,0)
        dfs(i,n-1)
    }
    for i := 0;i<n;i++{
        dfs(0,i)
        dfs(m-1,i)
    }
    for i , row := range grid{
        for j := range row{
            if grid[i][j] == 0{
                ans ++
                queue := make([][]int , 0)
                queue = append(queue , []int{i,j})
                for len(queue) > 0{
                    q := queue[0]
                    queue = queue[1:]
                    grid[q[0]][q[1]] = 2
                    for _ , d := range direction{
                        new_i , new_j := q[0] + d[0] , q[1] + d[1]
                        if (0 <= new_i && new_i < m && 0<=new_j && new_j < n) && grid[new_i][new_j] == 0 {
                            queue = append(queue , []int{new_i,new_j})
                        }
                    }
                }
            }
        }
    }
    return ans
}