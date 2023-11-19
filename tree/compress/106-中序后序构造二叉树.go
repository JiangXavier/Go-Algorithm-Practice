/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 || len(inorder) == 0 {
		return nil
	}
	num := postorder[len(postorder)-1]
	var id int
	for i, v := range inorder {
		if v == num {
			id = i
			break
		}
	}
	mm := len(inorder) - id - 1
	dd := len(postorder) - 1
	return &TreeNode{num, buildTree(inorder[:id], postorder[:dd-mm]), buildTree(inorder[id+1:], postorder[dd-mm:dd])}
}