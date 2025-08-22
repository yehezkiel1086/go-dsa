package main

import (
	"fmt"
)

/**
Traversal
Remove a node
Insert a node
Sort
*/

type t_node interface {
	traverse()
	findLowest() *Node
	deleteNode(*Node, *Node) *Node
}

type Node struct {
	data int
	next *Node
}

func (node *Node) traverse() {
	for curr := node; curr != nil; curr = curr.next {
		fmt.Print(curr.data, " -> ")
	}
	fmt.Println("null")
}

func (node *Node) findLowest() (*Node) {
	var lowest *Node
	lowest = node
	for tmp := node; tmp.next != nil; tmp = tmp.next {
		if tmp.data < lowest.data {
			lowest = tmp
		}
	}
	return lowest
}

func (node *Node) deleteNode(head *Node, nodeToDelete *Node) *Node {
	if head == nodeToDelete {
		return head.next
	}

	currNode := head
	for (currNode.next != nil && currNode.next != nodeToDelete) {
		currNode = currNode.next
	}

	if currNode.next == nil {
		return head
	}

	currNode.next = currNode.next.next

	return head
}

func main() {
	head := &Node{data: 5}
	head = &Node{data: 2, next: head}
	head = &Node{data: 4, next: head}
	head = &Node{data: 1, next: head}
	head = &Node{data: 3, next: head}

	var list t_node = head

	list.traverse()
	fmt.Println(list.findLowest())

	list = list.deleteNode(head, head.next)
	list.traverse()
}
