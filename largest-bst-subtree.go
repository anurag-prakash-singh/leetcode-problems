package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max2(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min2(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func bstSum(root *TreeNode, maxBSTSize *int) (bool, int, int, int) {
	if root.Left == nil && root.Right == nil {
		*maxBSTSize = max2(*maxBSTSize, 1)

		return true, 1, root.Val, root.Val
	}

	isTreeBST := false
	treeMin, treeMax := 0, 0
	treeSize := 0

	if root.Left != nil && root.Right != nil {
		isLSubTreeBST, lSubTreeSize, lSubTreeMin, lSubTreeMax := bstSum(root.Left, maxBSTSize)
		isRSubTreeBST, rSubTreeSize, rSubTreeMin, rSubTreeMax := bstSum(root.Right, maxBSTSize)

		isTreeBST = isLSubTreeBST && isRSubTreeBST && (root.Val > lSubTreeMax) && (root.Val < rSubTreeMin)
		treeSize = 1 + lSubTreeSize + rSubTreeSize
		treeMin, treeMax = min2(root.Val, min2(lSubTreeMin, rSubTreeMin)), max2(root.Val, max2(lSubTreeMax, rSubTreeMax))
	} else if root.Left == nil {
		isRSubTreeBST, rSubTreeSize, rSubTreeMin, rSubTreeMax := bstSum(root.Right, maxBSTSize)
		isTreeBST = isRSubTreeBST && (root.Val < rSubTreeMin)
		treeSize = 1 + rSubTreeSize
		treeMin, treeMax = min2(root.Val, rSubTreeMin), max2(root.Val, rSubTreeMax)
	} else {
		isLSubTreeBST, lSubTreeSize, lSubTreeMin, lSubTreeMax := bstSum(root.Left, maxBSTSize)
		isTreeBST = isLSubTreeBST && (root.Val > lSubTreeMax)
		treeSize = 1 + lSubTreeSize
		treeMin, treeMax = min2(root.Val, lSubTreeMin), max2(root.Val, lSubTreeMax)
	}

	if root.Val == -2 {
		fmt.Printf("treeMin: %v, treeMax: %v\n", treeMin, treeMax)
	}

	if isTreeBST {
		*maxBSTSize = max2(*maxBSTSize, treeSize)
	}

	return isTreeBST, treeSize, treeMin, treeMax
}

func largestBSTSubtree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	maxBSTSum := 0

	bstSum(root, &maxBSTSum)

	return maxBSTSum
}

func test1() {
	nodeM1 := TreeNode{Val: -1}
	nodeM2 := TreeNode{Val: -2}

	nodeM2.Right = &nodeM1

	result := largestBSTSubtree(&nodeM2)

	fmt.Printf("result: %d\n", result)
}

func main() {
	test1()
}
