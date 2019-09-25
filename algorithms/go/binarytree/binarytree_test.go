package binarytree

import (
	"reflect"
	"testing"
)

func TestBinaryTreeOrder(t *testing.T) {
	tree := &Node{Value: 1}
	tree.Lchild = &Node{Value: 2}
	tree.Rchild = &Node{Value: 5}

	// PreOrder
	var a []int
	tree.PreOrder(&a)
	if !reflect.DeepEqual(a, []int{1, 2, 5}) {
		t.Errorf("Got Wrong Order: %v", a)
	}

	// InOrder
	var b []int
	tree.InOrder(&b)
	if !reflect.DeepEqual(b, []int{2, 1, 5}) {
		t.Errorf("Got Wrong Order: %v", b)
	}

	// PostOrder
	var c []int
	tree.PostOrder(&c)
	if !reflect.DeepEqual(c, []int{2, 5, 1}) {
		t.Errorf("Got Wrong Order: %v", c)
	}
}
