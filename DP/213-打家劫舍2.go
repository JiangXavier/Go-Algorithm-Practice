//递推
//和打家劫舍1不同的点在于首位循环，则分为两种情况分别转化为第一类问题，考虑取不取第一个元素即可
func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}
	return max(s(nums[2:n-1])+nums[0], s(nums[1:]))
}
func s(nums []int) int {
	dp := make([]int, len(nums)+2)
	for i := 0; i < len(nums); i++ {
		dp[i+2] = max(dp[i+1], dp[i]+nums[i])
	}
	return dp[len(nums)+1]
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
    if n == 1 {return nums[0]}
    if n == 2 {return max(nums[0] , nums[1])}
    memo1 := make([]int , n + 1)
    memo2 := make([]int , n + 1)
    for i , _ := range memo1{
        memo1[i] = -1
        memo2[i] = -1
    }
    var dfs func(int , int , []int) int 
    dfs = func(begin int , i int, memo []int) int { 
        if begin > i {return 0}
        if memo[i] != -1 {return memo[i]}
        memo[i] = max( dfs(begin , i-2 , memo) + nums[i],dfs(begin , i - 1,memo))
        return memo[i]
    }
    return max(dfs(2,n-2,memo1) + nums[0] , dfs(1,n-1 , memo2))
}

func max(a,b int) int {
    if a > b {return a}
    return b
}