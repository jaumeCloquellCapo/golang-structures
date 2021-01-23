package structures

type TreeInterface interface {
	Insert(value int)
}

type BinaryTree struct {
	head *TNode
}


type TNode struct {
	left *TNode
	right *TNode
	value int
}

func NewBinaryTree() *BinaryTree {
	return &BinaryTree{
		head: nil,
	}
}

func NewTNode(value int, left *TNode,  right *TNode) *TNode {
	return &TNode{
		value: value,
		left: left,
		right: right,

	}
}

func (bt *BinaryTree) isEmpty() bool {
	return bt.head == nil
}


func (bt *BinaryTree) Insert(value int) {
	var temporalNode = NewTNode(value, nil, nil)

	if bt.isEmpty() {
		bt.head = temporalNode
	} else {
		var current = bt.head
		var parent *TNode

		for {
			parent = current
			//go to left of the tree
			if value < parent.value  {
				current = current.left
				//insert to the left
				if current == nil {
					parent.left = temporalNode
					break
				}
			} else {
				current = current.right
				//insert to the right
				if current == nil {
					parent.right = temporalNode
					break
				}
			}

		}
	}
}
