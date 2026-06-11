/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
    // find midpt and split into two
	fast := head; slow := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	// reverse second half
	prev := (*ListNode)(nil); curr := slow.Next
	slow.Next = nil
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr; curr = next
	}
	// merge two lists
	l1, l2 := head, prev
	for l2 != nil {
		temp := l1.Next; l1.Next = l2
		l1 = temp
		temp = l2.Next; l2.Next = l1
		l2 = temp
	}
}
