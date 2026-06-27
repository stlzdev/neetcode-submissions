/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p == nil && q == nil {
		return true
	}
	return check(p, q)
}

func check(a *TreeNode, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	} else if a == nil || b == nil {
		return false 
	} else if a.Val != b.Val {
		return false
	}
	return check(a.Left, b.Left) && check(a.Right, b.Right)
}
