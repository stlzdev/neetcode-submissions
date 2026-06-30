/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
    curr := root
    for curr != nil {
        if curr.Val == p.Val {
            return p
        }
        if curr.Val == q.Val {
            return q
        }
        if (p.Val < curr.Val && q.Val > curr.Val) || (p.Val > curr.Val && q.Val < curr.Val) {
            return curr 
        } else if p.Val > curr.Val && q.Val > curr.Val {
            curr = curr.Right
        } else {
            curr = curr.Left
        }
    }
    return root
}
