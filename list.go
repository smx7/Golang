package main

import (
	"fmt"
)

type ListNode struct {
	val  int
	next *ListNode
}

func CreateNode(value int) *ListNode {
	return &ListNode{value, nil}
}

func (list *ListNode) Print() {
	if list != nil {
		fmt.Println(list.val)
		list.next.Print()
	}
}

func (list *ListNode) PushBack(value int) {
	if list == nil {
		list = CreateNode(value)
		return
	}
	var prev *ListNode = list
	for list != nil {
		prev = list
		list = list.next
	}
	prev.next = CreateNode(value)
}

func (list *ListNode) PopBack() {
	if list == nil {
		return
	}
	var prev *ListNode = list
	for list.next != nil {
		prev = list
		list = list.next
	}
	prev.next = nil
}

func main() {
	//var list *ListNode = &ListNode{1, nil}
	//list.next = &ListNode{2, nil}
	var list ListNode
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)
	list.PushBack(4)
	list.PushBack(5)
	list.PopBack()
	list.Print()

}
