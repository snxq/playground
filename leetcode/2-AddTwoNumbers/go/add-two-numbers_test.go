package addtwonumbers

import (
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	newList := func(l []int) *ListNode {
		head := &ListNode{}
		pointer := head
		for _, i := range l {
			pointer.Val = i
			pointer.Next = &ListNode{}
			pointer = pointer.Next
		}
		return head
	}

	l1 := newList([]int{2, 4, 3})
	l2 := newList([]int{5, 6, 4})
	l3 := addTwoNumbers(l1, l2)

	p := l3
	for _, n := range []int{7, 0, 8} {
		if p.Val != n {
			t.Errorf("Result should be %d, but is %d", n, p.Val)
		}
		p = p.Next
	}
}
