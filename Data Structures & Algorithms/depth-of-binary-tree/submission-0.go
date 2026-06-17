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
	return addDepth(root, 0)
}

func addDepth(tree *TreeNode, curr int) int {
	if tree == nil {
		return curr
	}
	curr += 1
	left, right := tree.Left, tree.Right
	llen, rlen := addDepth(left, curr), addDepth(right, curr)
	return max(llen, rlen)
}
