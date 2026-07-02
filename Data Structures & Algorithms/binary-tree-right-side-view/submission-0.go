/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }
    out := []int{}
    queue := list.New()
    queue.PushBack(root)
    for queue.Len() > 0 {
        level := []int{}
        llen := queue.Len()
        for i := 0; i < llen; i++ {
            front := queue.Front()
            el := front.Value.(*TreeNode)
            queue.Remove(front)
            level = append(level, el.Val)
            if el.Left != nil {
                queue.PushBack(el.Left)
            }
            if el.Right != nil {
                queue.PushBack(el.Right)
            }
        }
        out = append(out, level[len(level) - 1])      
    }
    return out
}
