/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func goodNodes(root *TreeNode) int {
    if root == nil {
		return 0
	}
	currCount := 1
	if root.Left != nil {
		currCount += count(root.Left, root.Val)
	}
	if root.Right != nil {
		currCount += count(root.Right, root.Val)
	}
	return currCount
}

func count(node *TreeNode, max int) int {
	if node == nil {
		return 0
	}
	if node.Val >= max {
		max = node.Val
		return 1 + count(node.Left, max) + count(node.Right, max)
	} else {
		return count(node.Left, max) + count(node.Right, max)
	}
}

