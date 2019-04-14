package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	stack := make([]*TreeNode, 0, 5)
	traversalResult := make([]int, 0, 5)

	if root == nil {
		return []int{}
	}

	currNode := root

	for {
		// Push the left branch onto the stack
		for ; currNode != nil; currNode = currNode.Left {
			if currNode != nil {
				stack = append(stack, currNode)
			}
		}

		if len(stack) == 0 {
			break
		}

		topNode := stack[len(stack)-1]
		traversalResult = append(traversalResult, topNode.Val)
		stack = stack[:len(stack)-1]

		if topNode.Right != nil {
			currNode = topNode.Right
		}
	}

	return traversalResult
}

func test1() {
	node1 := TreeNode{1, nil, nil}
	node2 := TreeNode{2, nil, nil}
	node3 := TreeNode{3, nil, nil}

	node2.Left = &node1
	node2.Right = &node3

	result := inorderTraversal(&node2)

	fmt.Printf("result: %v\n", result)
}

func test2() {
	node1 := TreeNode{1, nil, nil}
	node2 := TreeNode{2, nil, nil}

	node2.Left = &node1

	result := inorderTraversal(&node2)

	fmt.Printf("result: %v\n", result)
}

func test3() {
	node1 := TreeNode{1, nil, nil}
	node3 := TreeNode{3, nil, nil}
	node2 := TreeNode{2, &node1, &node3}
	node5 := TreeNode{5, nil, nil}
	node6 := TreeNode{6, &node5, nil}
	node4 := TreeNode{4, &node2, &node6}

	result := inorderTraversal(&node4)

	fmt.Printf("result: %v\n", result)
}

func test4() {
	node1 := TreeNode{1, nil, nil}

	result := inorderTraversal(&node1)

	fmt.Printf("result: %v\n", result)
}

func main() {
	test1()
	test2()
	test3()
	test4()
}
