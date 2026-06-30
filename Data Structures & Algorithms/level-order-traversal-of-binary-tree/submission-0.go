/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
    queue := list.New()
    out := [][]int{}
    if root == nil {
        return nil
    }
    queue.PushBack(root)
    for queue.Len() > 0 {
        level := []int{}
        lsize := queue.Len()
        for i := 0; i < lsize; i++ {
            node := queue.Front()
            queue.Remove(node)
            el := node.Value.(*TreeNode)
            level = append(level, el.Val)
            if el.Left != nil {
                queue.PushBack(el.Left)
            }
            if el.Right != nil {
                queue.PushBack(el.Right)
            }
        }
        out = append(out, level)
    }
    return out
}
