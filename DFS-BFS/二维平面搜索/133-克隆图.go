/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
	v := map[*Node]*Node{}
	var dfs func(node *Node) *Node
	dfs = func(node *Node) *Node {
		if node == nil {
			return node
		}
		_, ok := v[node]
		if ok {
			return v[node]
		}
		clone_node := &Node{node.Val, []*Node{}}
		v[node] = clone_node
		for _, i := range node.Neighbors {
			clone_node.Neighbors = append(clone_node.Neighbors, dfs(i))
		}
		return clone_node
	}
	return dfs(node)
}