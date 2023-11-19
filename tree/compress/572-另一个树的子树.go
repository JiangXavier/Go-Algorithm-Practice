/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	q := []*TreeNode{}
	q = append(q, root)
	for len(q) != 0 {
		m := q[0]
		q = q[1:]
		if isSameTree(m, subRoot) {
			return true
		}
		if m.Left != nil {
			q = append(q, m.Left)
		}
		if m.Right != nil {
			q = append(q, m.Right)
		}
	}
	return false
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q != nil {
		return false
	}
	if q == nil && p != nil {
		return false
	}
	if p == nil && q == nil {
		return true
	}
	if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}