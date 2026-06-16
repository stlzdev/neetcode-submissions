/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    carry := 0
    dummy := &ListNode{Next: l1}
    tail := dummy
    for l1 != nil || l2 != nil {
        if l1 != nil && l2 != nil {
            sum := l1.Val + l2.Val + carry
            l1.Val = sum % 10
            carry = sum / 10
            l1 = l1.Next
            l2 = l2.Next
        } else if l1 != nil {
            sum := l1.Val + carry
            l1.Val = sum % 10
            carry = sum / 10
            l1 = l1.Next
        } else {
            sum := l2.Val + carry
            tail.Next = &ListNode{Val: sum % 10}
            carry = sum / 10
            l2 = l2.Next
        }
        tail = tail.Next
    }
    if carry > 0 {
        tail.Next = &ListNode{Val: carry}
    }
    return dummy.Next
}
