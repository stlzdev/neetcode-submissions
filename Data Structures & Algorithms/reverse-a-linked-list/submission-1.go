/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// Recursive Approach

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := (*ListNode)(head.Next)
	if head.Next != nil {
		newHead = reverseList(head.Next)
		head.Next.Next = head
	}
	head.Next = nil
	return newHead
}