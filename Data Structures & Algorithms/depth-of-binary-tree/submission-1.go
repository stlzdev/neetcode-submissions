/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	level := 0
	dq := list.New()
	dq.PushBack(root)
	for dq.Len() > 0 {
		for range dq.Len() {
			node := dq.Remove(dq.Front()).(*TreeNode)
			if node.Left != nil {
				dq.PushBack(node.Left)
			}
			if node.Right != nil {
				dq.PushBack(node.Right)
			}
		}
		level += 1
	}
	return level
}
