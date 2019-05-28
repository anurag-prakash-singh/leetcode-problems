package main

import (
	"errors"
	"fmt"
)

var errInvalidInput error = errors.New("Invalid input")

type node struct {
	data        int
	left, right *node
}

func addToTree(root *node, val int) *node {
	if root == nil {
		return &node{data: val}
	}

	if val <= root.data {
		root.left = addToTree(root.left, val)
	} else {
		root.right = addToTree(root.right, val)
	}

	return root
}

func getRightMostNode(root *node) *node {
	if root == nil {
		return nil
	}

	if root.right == nil {
		return root
	}

	return getRightMostNode(root.right)
}

func secondLargest(root *node) (int, error) {
	if root == nil || (root.left == nil && root.right == nil) {
		return 0, errInvalidInput
	}

	currNode := root

	for currNode != nil {
		if currNode.left != nil && currNode.right == nil {
			return getRightMostNode(currNode.left).data, nil
		}

		if currNode.right != nil && currNode.right.left == nil && currNode.right.right == nil {
			return currNode.data, nil
		}

		currNode = currNode.right
	}

	return -1, errInvalidInput
}

func buildTreeWithNums(nums []int) *node {
	var root *node

	for _, num := range nums {
		root = addToTree(root, num)
	}

	return root
}

func dumpTreeInorder(root *node) {
	if root == nil {
		return
	}

	dumpTreeInorder(root.left)
	fmt.Printf("%d ", root.data)
	dumpTreeInorder(root.right)
}

func checkSlicesEqual(slc1, slc2 []int) bool {
	if len(slc1) != len(slc2) {
		return false
	}

	for i := 0; i < len(slc1); i++ {
		if slc1[i] != slc2[i] {
			return false
		}
	}

	return true
}

func testBuildTree() {
	numss := [][]int{
		[]int{5, 4, 2, 1, 3},
		[]int{1, 3, 4, 2, 5},
		[]int{1, 2, 3, 4, 5},
		[]int{1, 2, 3, 5, 4},
	}

	for _, nums := range numss {
		root1 := buildTreeWithNums(nums)

		dumpTreeInorder(root1)

		fmt.Println()
	}
}

func testSecondLargest() {
	numss := [][]int{
		[]int{5, 4, 2, 1, 3},
		[]int{1, 3, 4, 2, 5},
		[]int{1, 2, 3, 4, 5},
		[]int{1, 2, 3, 5, 4},
		[]int{5, 4, 3, 2, 1},
	}

	for i, nums := range numss {
		root := buildTreeWithNums(nums)
		result, err := secondLargest(root)

		if err != nil {
			fmt.Printf("Error occurred: %v\n", err)

			continue
		}

		if result != 4 {
			fmt.Printf("FAILED test %d: Got %d, expected %d\n", i, result, 4)
		} else {
			fmt.Printf("PASSED test %d\n", i)
		}
	}
}

func main() {
	// testBuildTree()
	testSecondLargest()
}
