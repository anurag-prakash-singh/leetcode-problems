package main

import (
	"fmt"
	"sort"
)

type Ints []int

func (ints Ints) Len() int {
	return len(ints)
}

func (ints Ints) Less(i, j int) bool {
	return ints[i] < ints[j]
}

func (ints Ints) Swap(i, j int) {
	ints[i], ints[j] = ints[j], ints[i]
}

func findCombinations(candidates []int, target int, selectedCandidates []int, results [][]int) [][]int {
	if target < 0 {
		return results
	}

	if target == 0 {
		sel := make([]int, len(selectedCandidates))
		copy(sel, selectedCandidates)
		results = append(results, sel)

		return results
	}

	for i := 0; i < len(candidates); i++ {
		if i > 0 && candidates[i] == candidates[i-1] {
			continue
		}

		selectedCandidates = append(selectedCandidates, candidates[i])
		results = findCombinations(candidates[i+1:], target-candidates[i], selectedCandidates, results)
		selectedCandidates = selectedCandidates[:len(selectedCandidates)-1]
	}

	return results
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Sort(Ints(candidates))

	return findCombinations(candidates, target, []int{}, [][]int{})
}

func tests() {
	fmt.Printf("%v\n", combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
	fmt.Printf("%v\n", combinationSum2([]int{2, 5, 2, 1, 2}, 5))
	fmt.Printf("%v\n", combinationSum2([]int{1, 1, 2, 2}, 5))
}

func main() {
	tests()
}
