package main

import (
	"fmt"
)

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func CreateNode(value int) *TreeNode {
	return &TreeNode{value, nil, nil}

}

func (node *TreeNode) InOrder() {
	if node != nil {
		node.Left.InOrder()
		fmt.Println(node.Value)
		node.Right.InOrder()
	}
}

func (node *TreeNode) Insert(value int) bool {
	if node == nil {
		node = CreateNode(value)
		return true
	}
	var parent *TreeNode = nil
	var cur *TreeNode = node
	for cur != nil {
		if cur.Value > value {
			parent = cur
			cur = cur.Left
		} else if cur.Value < value {
			parent = cur
			cur = cur.Right
		} else {
			return false
		}
	}
	cur = CreateNode(value)
	if cur.Value < parent.Value {
		parent.Left = cur
	} else {
		parent.Right = cur
	}
	return true
}

func (node *TreeNode) Erase(value int) bool {
	if node == nil {
		return false
	}
	if node.Value == value {
		node = nil
		return true
	}
	var parent *TreeNode
	var cur *TreeNode = node
	for cur != nil {
		if cur.Value > value {
			parent = cur
			cur = cur.Left
		} else if cur.Value < value {
			parent = cur
			cur = cur.Right
		} else {
			if cur.Left == nil {
				if cur == parent.Left {
					parent.Left = cur.Right
				} else {
					parent.Right = cur.Right
				}
			} else if cur.Right == nil {
				if cur == parent.Left {
					parent.Left = cur.Left
				} else {
					parent.Right = cur.Right
				}
			} else {
				var RightMinLeft *TreeNode = cur.Right
				for RightMinLeft.Left != nil {
					parent = RightMinLeft
					RightMinLeft = RightMinLeft.Left
				}
				cur.Value = RightMinLeft.Value
				if RightMinLeft == parent.Left {
					parent.Left = RightMinLeft.Right
				} else {
					parent.Right = RightMinLeft.Right
				}
				RightMinLeft = nil
			}
			return true
		}
	}
	return false
}
func main() {
	var root TreeNode
	root.Insert(4)
	root.Erase(4)
	fmt.Println("******************")
	//root.Value = 4
	root.Insert(2)
	root.Insert(6)
	//root.Insert(4)
	root.Insert(3)
	root.Insert(5)
	root.Insert(1)
	root.Insert(7)
	ret := root.Erase(3)
	fmt.Println(ret)
	root.InOrder()
}
