/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	_, val := f(nums)
	return &TreeNode{nums[val], constructMaximumBinaryTree(nums[:val]), constructMaximumBinaryTree(nums[val+1:])}
}
func f(nums []int) (i, j int) {
	max := -1
	wei := 0
	for m, n := range nums {
		if n > max {
			max = n
			wei = m
		}
	}
	return max, wei
}

// 递归构建树，直接返回对应树结构，左右子树部分分别递归函数，每一次函数调用为一个节点赋值