package main

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

func max3(a, b, c int) int {
	return max2(a, max2(b, c))
}

func traverse(root *TreeNode, maxSum *int) int {
	if root == nil {
		return 0
	}

	lPathMaxSum := traverse(root.Left, maxSum)
	rPathMaxSum := traverse(root.Right, maxSum)
	rootPathMaxSum := max3(root.Val, root.Val+lPathMaxSum, root.Val+rPathMaxSum)

	*maxSum = max3(*maxSum, root.Val+lPathMaxSum+rPathMaxSum, rootPathMaxSum)

	return rootPathMaxSum
}

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	maxSum := root.Val
	traverse(root, &maxSum)

	return maxSum
}

func main() {

}
