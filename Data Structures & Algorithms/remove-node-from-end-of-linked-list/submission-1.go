/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    if head == nil {
		return head
	}
	dummy := &ListNode{Next: head}
	first := head
	second := dummy
	for ahead := 0; ahead < n; ahead++ {
		first = first.Next
	}
	if first == nil {
		return head.Next
	}
	for first != nil {
		first = first.Next
		second = second.Next
	}
	second.Next = second.Next.Next
	return dummy.Next
}
