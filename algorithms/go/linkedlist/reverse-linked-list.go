package linkedlist

// ReverseListNode 反转单链表
func ReverseListNode(head *SingleListNode) *SingleListNode {
	var pre, next *SingleListNode

	for head != nil {
		next = head.Next
		head.Next = pre
		pre = head
		head = next
	}

	return pre
}
