package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func allPossibleFBT(N int) []*TreeNode {
	if N%2 == 0 {
		return []*TreeNode{}
	}

	if N == 1 {
		return []*TreeNode{&(TreeNode{Val: 0, Left: nil, Right: nil})}
	}

	possibleTrees := make([]*TreeNode, 0, 5)

	for i := 1; i < N-1; i++ {
		lN := i
		rN := N - 1 - i

		lTrees := allPossibleFBT(lN)
		rTrees := allPossibleFBT(rN)

		for _, lTree := range lTrees {
			for _, rTree := range rTrees {
				root := TreeNode{Val: 0, Left: lTree, Right: rTree}
				possibleTrees = append(possibleTrees, &root)
			}
		}
	}

	return possibleTrees
}
