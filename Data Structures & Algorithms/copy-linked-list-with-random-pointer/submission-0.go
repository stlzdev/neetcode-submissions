/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    hmap := make(map[*Node] *Node)
	curr := head
	for curr != nil {
		hmap[curr] = &Node{Val: curr.Val}
		curr = curr.Next
	}
	curr = head
	for curr != nil {
		hmap[curr].Next = hmap[curr.Next]
		hmap[curr].Random = hmap[curr.Random]
		curr = curr.Next
	}
	return hmap[head]
}
