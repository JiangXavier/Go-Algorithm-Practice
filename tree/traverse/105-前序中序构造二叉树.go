/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	num := preorder[0]
	var id int
	for i, v := range inorder {
		if v == num {
			id = i
			break
		}
	}
	return &TreeNode{num, buildTree(preorder[1:1+id], inorder[:id]), buildTree(preorder[1+id:], inorder[id+1:])}
}