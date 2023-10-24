//排序，nlogn
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    nums := append(nums1 , nums2...)
    sort.Ints(nums)
    if len(nums) % 2 == 1{
        return float64(nums[len(nums) / 2])
    } else{
        return float64(nums[len(nums) / 2] + nums[(len(nums) / 2) - 1]) / 2.0
    }
}
