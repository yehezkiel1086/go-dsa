package main

import "fmt"

/**
Traversal
Find lowest value
Remove a node
Insert a node
Sort
*/

type Node struct {
	Data int
	Next *Node
}

func (n *Node) traverse() {
	for n != nil {
		fmt.Printf("%v -> ", n.Data)
		n = n.Next
	}
	fmt.Println("nil")
}

func (n *Node) findLowest() int {
	lowest := n.Data
	n = n.Next

	for n != nil {
		if n.Data < lowest {
			lowest = n.Data
		}
		n = n.Next
	}

	return lowest
}

func (head *Node) deleteNode(node *Node) *Node {
	// if node to delete is head
	if head == node {
		return head.Next
	}

	// find node to delete
	curr := head
	for curr.Next != nil && curr.Next != node {
		curr = curr.Next
	}

	// if no node to delete
	if curr.Next == nil {
		return head
	}

	// if node to delete is found
	curr.Next = curr.Next.Next

	return head
}

func (head *Node) insert(node *Node, pos int) *Node {
	// if position == 1
	if pos == 1 {
		node.Next = head
		return node
	}

	// traverse
	curr := head
	for i := 0; i < pos - 1 && curr != nil; i++ {
		curr = curr.Next
	}

	if curr == nil {
		return head
	}

	// assign new node
	node.Next = curr.Next
	curr.Next = node
	return head
}

func main() {
	node5 := &Node{
		Data: 9,
	}
	node4 := &Node{
		Data: 2,
		Next: node5,
	}
	node3 := &Node{
		Data: 3,
		Next: node4,
	}
	node2 := &Node{
		Data: 11,
		Next: node3,
	}
	head := &Node{
		Data: 7,
		Next: node2,
	}

	// traverse
	head.traverse()

	// find lowest
	lowest := head.findLowest()
	fmt.Println(lowest)

	// delete node
	head = head.deleteNode(node3)
	head.traverse()

	// insert node
	newNode := &Node{
		Data: 10,
	}
	head = head.insert(newNode, 3)
	head.traverse()
}
