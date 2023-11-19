/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	if len(preorder) == 0 || len(postorder) == 0 {
		return nil
	}
	num := preorder[0]
	if len(preorder) == 1 {
		return &TreeNode{Val: num}
	}
	id := 0
	for i, j := range postorder {
		if j == preorder[1] {
			id = i
			break
		}
	}
	return &TreeNode{num, constructFromPrePost(preorder[1:2+id], postorder[:id+1]), constructFromPrePost(preorder[2+id:], postorder[id+1:len(postorder)-1])}
}