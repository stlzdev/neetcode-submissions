/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
    if node == nil {
		return nil
	}
	hmap := make(map[*Node]*Node)
	var dfs func(*Node) *Node
	dfs = func(curr *Node) *Node {
		if cloned, exists := hmap[curr]; exists {
			return cloned
		}
		clone := &Node{Val: curr.Val}
		hmap[curr] = clone
		for _, neighbor := range curr.Neighbors {
			clone.Neighbors = append(clone.Neighbors, dfs(neighbor))
		}
		return clone
	}
	return dfs(node)
}
