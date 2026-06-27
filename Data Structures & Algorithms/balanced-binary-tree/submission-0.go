/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
    if root == nil {
		return true
	}
	return balanced(root) != -1
}

func balanced(node *TreeNode) int {
	if node == nil {
		return 0
	}
	ll, rl := balanced(node.Left), balanced(node.Right)
	if ll == -1 || rl == -1 || max(ll, rl) - min(ll, rl) > 1 {
		return -1
	}
	return 1 + max(ll, rl)
}
