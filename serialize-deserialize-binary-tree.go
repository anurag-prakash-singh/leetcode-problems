package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	left, right *TreeNode
	data        int
}

func serialize(root *TreeNode) string {
	sxd := []string{}
	stack := []*TreeNode{}

	if root == nil {
		return ""
	}

	stack = append(stack, root)

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if top == nil {
			sxd = append(sxd, "#")

			continue
		}

		sxd = append(sxd, strconv.Itoa(top.data))
		stack = append(stack, top.right)
		stack = append(stack, top.left)
	}

	return strings.Join(sxd, ",")
}

func deserializeInternal(treeStrParts []string, currPos *int) *TreeNode {
	if len(treeStrParts) == 0 {
		return nil
	}

	if treeStrParts[*currPos] == "#" {
		*currPos++

		return nil
	}

	nodeData, err := strconv.Atoi(treeStrParts[*currPos])

	if err != nil {
		fmt.Printf("Invalid data: %v\n", treeStrParts[*currPos])

		return nil
	}

	newNode := &TreeNode{data: nodeData}
	*currPos++
	newNode.left = deserializeInternal(treeStrParts, currPos)
	newNode.right = deserializeInternal(treeStrParts, currPos)

	return newNode
}

func deserialize(treeStr string) *TreeNode {
	pos := 0

	return deserializeInternal(strings.Split(treeStr, ","), &pos)
}

func test1() {
	node2 := TreeNode{data: 2}
	node4 := TreeNode{data: 4}
	node5 := TreeNode{data: 5}
	node3 := TreeNode{data: 3, left: &node4, right: &node5}
	node1 := TreeNode{data: 1, left: &node2, right: &node3}

	sxdResult := serialize(&node1)

	fmt.Printf("Serialized form: %v\n", sxdResult)

	dsxd := deserialize(sxdResult)

	fmt.Printf("Serialized form of deserialized tree: %v\n", serialize(dsxd))
}

func main() {
	fmt.Printf("Serialize/deserialize binary tree\n")

	test1()
}
