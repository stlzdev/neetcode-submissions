/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func invertTree(root *TreeNode) *TreeNode {
    if root == nil {
		return nil
	}
	leftInv := invertTree(root.Left)
	rightInv := invertTree(root.Right)
	root.Left = rightInv
	root.Right = leftInv
	return root
}
