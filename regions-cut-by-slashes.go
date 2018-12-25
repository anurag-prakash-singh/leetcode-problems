package main

import "fmt"

type UnionFind interface {
	// Union items with IDs id1 and id2
	Union(id1, id2 int)

	// Returns true if items with IDs id1 and id2 are in the same set
	Find(id1, id2 int) bool

	// Returns the ID of the root element of the set that the element
	// identified by `id` is part of
	Root(id int) int
}

type GridElems []int

func (g GridElems) Root(id int) int {
	for id != g[id] {
		id = g[id]
	}

	return id
}

func (g GridElems) Find(id1, id2 int) bool {
	return g.Root(id1) == g.Root(id2)
}

func (g GridElems) Union(id1, id2 int) {
	rootId1 := g.Root(id1)
	rootId2 := g.Root(id2)

	g[rootId1] = rootId2
}

func initializeIDs(ids []int) {
	for i := 0; i < len(ids); i++ {
		ids[i] = i
	}
}

func regionsBySlashes(grid []string) int {
	n := len(grid)
	regions := GridElems(make([]int, n*n*4))

	initializeIDs(regions)

	gridChars := make([][]rune, n)

	for i := 0; i < n; i++ {
		gridChars[i] = []rune(grid[i])
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			cellNum := i*n + j
			uQ := cellNum * 4
			rQ := cellNum*4 + 1
			dQ := cellNum*4 + 2
			lQ := cellNum*4 + 3

			// merge with left rQ of cell
			if j > 0 {
				lCellNum := i*n + (j - 1)
				rlQ := lCellNum*4 + 1
				regions.Union(lQ, rlQ)
			}

			// merge with up dQ of cell
			if i > 0 {
				uCellNum := (i-1)*n + j
				duQ := uCellNum*4 + 2
				regions.Union(uQ, duQ)
			}

			// merge with right lQ of cell
			if j < n-1 {
				rCellNum := i*n + (j + 1)
				lrQ := rCellNum*4 + 3
				regions.Union(rQ, lrQ)
			}

			// merge with down uQ of cell
			if i < n-1 {
				dCellNum := (i+1)*n + j
				udQ := dCellNum * 4
				regions.Union(dQ, udQ)
			}

			switch gridChars[i][j] {
			case ' ':
				regions.Union(dQ, uQ)
				regions.Union(lQ, rQ)
				regions.Union(dQ, rQ)
			case '\\':
				regions.Union(dQ, lQ)
				regions.Union(rQ, uQ)
			case '/':
				regions.Union(dQ, rQ)
				regions.Union(lQ, uQ)
			}
		}
	}

	// Count unique regions
	uniqueRegions := make([]int, len(regions))

	for _, r := range regions {
		regionRoot := regions.Root(r)

		uniqueRegions[regionRoot] = 1
	}

	uniqueRegionCount := 0

	for _, s := range uniqueRegions {
		if s == 0 {
			continue
		}

		uniqueRegionCount++
	}

	return uniqueRegionCount
}

func test1() {
	grid := []string{
		"//",
		"/ ",
	}

	fmt.Printf("%d\n", regionsBySlashes(grid))
}

func test2() {
	grid := []string{
		"/\\",
		"\\/",
	}

	fmt.Printf("%d\n", regionsBySlashes(grid))
}

func test3() {
	grid := []string{
		"\\/",
		"/\\",
	}

	fmt.Printf("%d\n", regionsBySlashes(grid))
}

func test4() {
	grid := []string{
		" /",
		"  ",
	}

	fmt.Printf("%d\n", regionsBySlashes(grid))
}

func test5() {
	grid := []string{
		" /",
		"/ ",
	}

	fmt.Printf("%d\n", regionsBySlashes(grid))
}

func main() {
	test1()
	test2()
	test3()
	test4()
	test5()
}
