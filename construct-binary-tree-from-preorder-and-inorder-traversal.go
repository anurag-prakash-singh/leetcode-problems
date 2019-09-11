package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getSubTree(preorder []int, preorderStart, preorderEnd int, inorder []int, inorderStart, inorderEnd int, preorderOffsetLUT, inorderOffsetLUT map[int]int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 || preorderStart > preorderEnd || inorderStart > inorderEnd {
		return nil
	}

	root := preorder[preorderStart]
	rootInorderOffset := inorderOffsetLUT[root]
	lSubTreeLen := rootInorderOffset - inorderStart

	if lSubTreeLen < 0 {
		fmt.Printf("****ERROR******\n")
		fmt.Printf("root: %d; preorderStart: %d; inorderStart: %d; rootInorderOffset: %d; lSubTreeLen: %d\n",
			root, preorderStart, inorderStart, rootInorderOffset, lSubTreeLen)

		return nil
	}

	// fmt.Printf("getting lSubTree\n")
	lSubTree := getSubTree(preorder, preorderStart+1, preorderStart+1+lSubTreeLen-1,
		inorder, inorderStart, inorderStart+lSubTreeLen-1, preorderOffsetLUT, inorderOffsetLUT)

	// fmt.Printf("getting rSubTree\n")
	rSubTree := getSubTree(preorder, preorderStart+lSubTreeLen+1, preorderEnd,
		inorder, rootInorderOffset+1, inorderEnd, preorderOffsetLUT, inorderOffsetLUT)

	return &TreeNode{Val: root, Left: lSubTree, Right: rSubTree}
}

func inorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}

	inorderTraversal(root.Left)
	fmt.Printf("%d ", root.Val)
	inorderTraversal(root.Right)
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	preorderOffsetLUT := make(map[int]int)
	inorderOffsetLUT := make(map[int]int)

	for i := 0; i < len(preorder); i++ {
		preorderOffsetLUT[preorder[i]] = i
		inorderOffsetLUT[inorder[i]] = i
	}

	return getSubTree(preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1, preorderOffsetLUT, inorderOffsetLUT)
}

func test1() {
	// preorder := []int{3, 9, 20, 15, 7}
	// inorder := []int{9, 3, 15, 20, 7}
	preorder := []int{5, 1, 9}
	inorder := []int{1, 5, 9}
	// preorder := []int{1, 5}
	// inorder := []int{1, 5}

	tree := buildTree(preorder, inorder)

	inorderTraversal(tree)
}

func main() {
	test1()
}
