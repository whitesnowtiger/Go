package main

import "fmt"

type Node struct {
	next *Node
	val  int
}

type LinkedList struct {
	root *Node
	tail *Node
}

func (l *LinkedList) AddNode(val int) *Node {
	if l.root == nil {
		l.root = &Node{val: val}
		l.tail = l.root
		return l.tail
	}
	l.tail.next = &Node{val: val}
	l.tail = l.tail.next
	return l.tail
}

func (l *LinkedList) RemoveNode(node *Node) (*Node, *Node) {
	if node == l.root {
		l.root = l.root.next
		if l.root == nil {
			l.tail = nil
		}
		return l.root, l.tail
	}

	prev := l.root
	for prev.next != node {
		prev = prev.next
	}

	if node == l.tail {
		prev.next = nil
		l.tail = prev
	} else {
		prev.next = prev.next.next
	}
	return l.root, l.tail
}

func (l *LinkedList) PrintNodes() {
	node := l.root
	for node.next != nil {
		fmt.Printf("%d -> ", node.val)
		node = node.next
	}
	fmt.Printf("%d\n", node.val)
}

func main() {
	list := &LinkedList{}
	list.AddNode(0)

	for i := 1; i < 10; i++ {
		list.AddNode(i)
	}

	list.PrintNodes()

	list.RemoveNode(list.root.next)

	list.PrintNodes()

	list.RemoveNode(list.root)

	list.PrintNodes()

	list.RemoveNode(list.tail)

	list.PrintNodes()

}
