package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func traverse(root *TreeNode, voyage []int, voyageOffset *int) []int {
	if *voyageOffset >= len(voyage) || root.Val != voyage[*voyageOffset] {
		return []int{-1}
	}

	if root.Left == nil && root.Right == nil {
		return []int{}
	}

	if root.Left != nil && root.Right != nil {
		if *voyageOffset >= len(voyage) {
			return []int{-1}
		}

		flips := []int{}

		if root.Left.Val != voyage[*voyageOffset+1] {
			root.Left, root.Right = root.Right, root.Left
			flips = append(flips, root.Val)
		}

		*voyageOffset++
		flipsL := traverse(root.Left, voyage, voyageOffset)

		if len(flipsL) > 0 && flipsL[0] == -1 {
			return []int{-1}
		}

		flips = append(flips, flipsL...)

		*voyageOffset++
		flipsR := traverse(root.Right, voyage, voyageOffset)

		if len(flipsR) > 0 && flipsR[0] == -1 {
			return []int{-1}
		}

		flips = append(flips, flipsR...)

		return flips
	} else if root.Left != nil {
		*voyageOffset++
		return traverse(root.Left, voyage, voyageOffset)
	} else {
		*voyageOffset++
		return traverse(root.Right, voyage, voyageOffset)
	}
}

func flipMatchVoyage(root *TreeNode, voyage []int) []int {
	voyageOffset := 0

	ans := traverse(root, voyage, &voyageOffset)

	// fmt.Printf("ans (tent): %v\n", ans)

	if len(ans) > 0 && ans[0] != -1 && voyageOffset < len(voyage)-1 {
		return []int{-1}
	}

	return ans
}

func modifyInt(val *int) {
	*val = 4
}

func test1() {
	root2 := TreeNode{Val: 2, Left: nil, Right: nil}
	root3 := TreeNode{Val: 3, Left: nil, Right: nil}
	root1 := TreeNode{Val: 1, Left: &root2, Right: &root3}

	voyage := []int{1, 3, 2}

	fmt.Printf("ans: %v\n", flipMatchVoyage(&root1, voyage))
}

func test2() {
	root2 := TreeNode{Val: 2, Left: nil, Right: nil}
	root3 := TreeNode{Val: 3, Left: nil, Right: nil}
	root1 := TreeNode{Val: 1, Left: &root2, Right: &root3}

	voyage := []int{1, 2, 3}

	fmt.Printf("ans: %v\n", flipMatchVoyage(&root1, voyage))
}

func test3() {
	root2 := TreeNode{Val: 2, Left: nil, Right: nil}
	root1 := TreeNode{Val: 1, Left: &root2, Right: nil}

	voyage := []int{2, 1}

	fmt.Printf("ans: %v\n", flipMatchVoyage(&root1, voyage))
}

func test4() {
	root2 := TreeNode{Val: 2, Left: nil, Right: nil}
	root4 := TreeNode{Val: 4, Left: nil, Right: nil}
	root5 := TreeNode{Val: 5, Left: nil, Right: nil}
	root3 := TreeNode{Val: 3, Left: &root4, Right: &root5}
	root1 := TreeNode{Val: 1, Left: &root2, Right: &root3}

	voyage := []int{1, 3, 4, 5, 2}

	fmt.Printf("ans: %v\n", flipMatchVoyage(&root1, voyage))
}

func test5() {
	root2 := TreeNode{Val: 2, Left: nil, Right: nil}
	root4 := TreeNode{Val: 4, Left: nil, Right: nil}
	root5 := TreeNode{Val: 5, Left: nil, Right: nil}
	root3 := TreeNode{Val: 3, Left: &root4, Right: &root5}
	root1 := TreeNode{Val: 1, Left: &root2, Right: &root3}

	voyage := []int{1, 3, 5, 4, 2}

	fmt.Printf("ans: %v\n", flipMatchVoyage(&root1, voyage))
}

func test6() {
	root2 := TreeNode{Val: 2, Left: nil, Right: nil}
	root4 := TreeNode{Val: 4, Left: nil, Right: nil}
	root5 := TreeNode{Val: 5, Left: nil, Right: nil}
	root3 := TreeNode{Val: 3, Left: &root4, Right: &root5}
	root1 := TreeNode{Val: 1, Left: &root2, Right: &root3}

	voyage := []int{1, 3, 5}

	fmt.Printf("ans: %v\n", flipMatchVoyage(&root1, voyage))
}

func test7() {
	root2 := TreeNode{Val: 2, Left: nil, Right: nil}
	root4 := TreeNode{Val: 4, Left: nil, Right: nil}
	root5 := TreeNode{Val: 5, Left: nil, Right: nil}
	root3 := TreeNode{Val: 3, Left: &root4, Right: &root5}
	root1 := TreeNode{Val: 1, Left: &root2, Right: &root3}

	voyage := []int{1, 3, 5, 4, 2, 7}

	fmt.Printf("ans: %v\n", flipMatchVoyage(&root1, voyage))
}

func main() {
	// a := 5

	// modifyInt(&a)

	// fmt.Printf("a = %d\n", a)

	test1()
	test2()
	test3()
	test4()
	test5()
	test6()
	test7()
}
