package main

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

func isIdentical(t1 *TreeNode, t2 *TreeNode) bool {
	if t1 == nil {
		return t2 == nil
	}

	if t2 == nil {
		return false
	}

	return t1.Val == t2.Val && isIdentical(t1.Left, t2.Left) && isIdentical(t1.Right, t2.Right)
}

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if t == nil {
		// bug in Leetcode's solution checked. A nil t should be
		// considered a subtree of any s.
		return s == nil
	}

	if isIdentical(s, t) {
		return true
	}

	if s == nil {
		return false
	}

	return isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func main() {

}
