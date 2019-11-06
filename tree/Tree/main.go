package main

import "fmt"

type TreeNode struct {
	Val    int
	Childs []*TreeNode
}

type Tree struct {
	Root *TreeNode
}

func (t *TreeNode) AddNode(val int) {
	t.Childs = append(t.Childs, &TreeNode{Val: val})
}

func (t *Tree) AddNode(val int) {
	if t.Root == nil {
		t.Root = &TreeNode{Val: val}
	} else {
		t.Root.Childs = append(t.Root.Childs, &TreeNode{Val: val})
	}
}

func DFS1(node *TreeNode) {
	fmt.Printf("%d-> ", node.Val)

	for i := 0; i < len(node.Childs); i++ {
		DFS1(node.Childs[i])
	}

}

func (t *Tree) DFS1() {
	DFS1(t.Root)
}

func (t *Tree) DFS2() {
	stack := []*TreeNode{}
	stack = append(stack, t.Root)

	for len(stack) > 0 {
		var last *TreeNode
		last, stack = stack[(len(stack)-1)], stack[:len(stack)-1]

		fmt.Printf("%d-> ", last.Val)

		for i := 0; i < len(last.Childs); i++ {
			stack = append(stack, last.Childs[i])
		}
	}
}

func (t *Tree) DFS3() {
	stack := []*TreeNode{}
	stack = append(stack, t.Root)

	for len(stack) > 0 {
		var last *TreeNode
		last, stack = stack[(len(stack)-1)], stack[:len(stack)-1]

		fmt.Printf("%d-> ", last.Val)

		for i := len(last.Childs) - 1; i >= 0; i-- {
			stack = append(stack, last.Childs[i])
		}
	}
}

func (t *Tree) BFS() {
	queue := []*TreeNode{}
	queue = append(queue, t.Root)

	for len(queue) > 0 {
		var front *TreeNode
		front, queue = queue[0], queue[1:]

		fmt.Printf("%d-> ", front.Val)

		for i := 0; i < len(front.Childs); i++ {
			queue = append(queue, front.Childs[i])

		}

	}
}

func main() {
	namu := Tree{}

	val := 1
	namu.AddNode(val)
	val++

	for i := 0; i < 3; i++ {
		namu.Root.AddNode(val)
		val++
	}

	for i := 0; i < len(namu.Root.Childs); i++ {
		for j := 0; j < 2; j++ {
			namu.Root.Childs[i].AddNode(val)
			val++
		}
	}
	namu.DFS1()

	fmt.Println()

	namu.DFS2()

	fmt.Println()

	namu.DFS3()

	fmt.Println()

	namu.BFS()

}
