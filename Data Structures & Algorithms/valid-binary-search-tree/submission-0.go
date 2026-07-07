/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	left := valid(root.Left, -99999, root.Val)
	right := valid(root.Right, root.Val, 99999)
	return left && right
}

func valid(node *TreeNode, min int, max int) bool {
	if node == nil {
		return true
	}
	if node.Val <= min || node.Val >= max {
		return false
	}
	left := true
	right := true
	if node.Left != nil {
		left = valid(node.Left, min, node.Val)
	}
	if node.Right != nil {
		right = valid(node.Right, node.Val, max)
	}
	return left && right
}