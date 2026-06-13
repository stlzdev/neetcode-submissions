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
	curr2 := head
	for curr2 != nil {
		hmap[curr2].Next = hmap[curr2.Next]
		hmap[curr2].Random = hmap[curr2.Random]
		curr2 = curr2.Next
	}
	return hmap[head]
}
