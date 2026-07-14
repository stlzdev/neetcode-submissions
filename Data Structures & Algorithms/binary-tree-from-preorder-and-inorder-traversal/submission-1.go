/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
    root := preorder[0]
	tree := &TreeNode{Val: root}
	idxMap := make(map[int]int, len(inorder))
	for idx, num := range inorder {
		idxMap[num] = idx
	}
	rootIdx := idxMap[root]
	tree.Left = buildTree(preorder[1:1+rootIdx], inorder[0:rootIdx])
	tree.Right = buildTree(preorder[1+rootIdx:], inorder[rootIdx+1:])
	return tree
}