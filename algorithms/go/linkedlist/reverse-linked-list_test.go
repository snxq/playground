package linkedlist

import "testing"

func TestReverseListNode(t *testing.T) {
	newSingleList := func(nodes []int) *SingleListNode {
		head := &SingleListNode{}
		pointer := head
		for _, n := range nodes {
			pointer.Next = &SingleListNode{Val: n}
			pointer = pointer.Next
		}
		return head.Next
	}

	cases := []struct {
		Expect *SingleListNode
		Actual *SingleListNode
	}{
		{
			Expect: newSingleList([]int{3, 2, 1}),
			Actual: ReverseListNode(newSingleList([]int{1, 2, 3})),
		},
		{
			Expect: newSingleList([]int{33, 22, 11}),
			Actual: ReverseListNode(newSingleList([]int{11, 22, 33})),
		},
	}

	for _, c := range cases {

		e := c.Expect
		a := c.Actual
		for e != nil || a != nil {
			println(a.Val.(int))
			if e.Val.(int) != a.Val.(int) {
				t.Errorf("Should be %v, but %v got.", e.Val, a.Val)
			}

			e = e.Next
			a = a.Next
		}
	}
}
