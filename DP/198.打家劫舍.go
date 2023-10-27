//暴力回溯--超时
func rob(nums []int) int {
	n := len(nums)
	var dfs func(int) int
	dfs = func(i int) int {
		if i < 0 {
			return 0
		}
		return max(dfs(i-1), dfs(i-2)+nums[i])
	}
	return dfs(n - 1)
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//记忆化搜索
func rob(nums []int) int {
    n := len(nums)
    memo := make([]int , n + 1)
    for i , _ := range memo{
        memo[i] = -1
    }
    var dfs func(int) int
    dfs = func(i int) int{
        if i < 0 {
            return 0
        }
        if memo[i] != -1 {return memo[i]}
        memo[i] = max(dfs(i-1) , dfs(i-2) + nums[i])
        return memo[i]
    }
    return dfs(n-1)
}
func max(a , b int) int{
    if a > b {return a}
    return b
}
//递推
func rob(nums []int) int {
    n := len(nums)
    dp := make([]int , n + 1)
    dp[0] = nums[0]
    if n < 2{return nums[0]}  
    dp[1] = max(nums[0] , nums[1])
    for i := 2 ; i < n ; i++{
        dp[i] = max(dp[i-1] , dp[i-2] + nums[i])
    }
    return dp[n-1]
}
func max(a , b int) int{
    if a > b {return a}
    return b
}
// 空间优化
func rob(nums []int) int {
    n := len(nums)
    if n < 2{return nums[0]} 
    a , b := nums[0] , max(nums[0] , nums[1])
    for i := 2 ; i < n ; i++{
        c := max(b , a + nums[i])
        a = b
        b = c
    }
    return b
}
func max(a , b int) int{
    if a > b {return a}
    return b
}