package binarytree

// Node binarytree node
type Node struct {
	Value          int
	Lchild, Rchild *Node
}

// PreOrder .
func (n *Node) PreOrder(arr *[]int) {
	if n == nil {
		return
	}
	*arr = append(*arr, n.Value)
	n.Lchild.PreOrder(arr)
	n.Rchild.PreOrder(arr)
}

// InOrder .
func (n *Node) InOrder(arr *[]int) {
	if n == nil {
		return
	}
	n.Lchild.InOrder(arr)
	*arr = append(*arr, n.Value)
	n.Rchild.InOrder(arr)
}

// PostOrder .
func (n *Node) PostOrder(arr *[]int) {
	if n == nil {
		return
	}
	n.Lchild.PostOrder(arr)
	n.Rchild.PostOrder(arr)
	*arr = append(*arr, n.Value)
}
