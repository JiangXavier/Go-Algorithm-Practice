func deleteAndEarn(nums []int) int {
    da := 0
    m := make([]int , 100001)
    for _ , x := range nums{
        m[x] ++
        da = max(da , x)
    }
    dp := make([]int , da + 1)
    dp[1] = m[1]
    for i := 2 ; i < da + 1;i++{
        dp[i] = max(dp[i-1] , dp[i-2] + i * m[i])
    }
    return dp[da]
}
func max(a , b int) int{
    if a > b {return a}
    return b
}