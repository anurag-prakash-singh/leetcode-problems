package main

import (
	"fmt"
	"sort"
)

type ByRune []rune

func (b ByRune) Len() int           { return len(b) }
func (b ByRune) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b ByRune) Less(i, j int) bool { return b[i] < b[j] }

func groupAnagrams(strs []string) [][]string {
	anagrams := make(map[string][]int)

	for i, s := range strs {
		runes := ByRune(s)
		sort.Sort(runes)
		normalized := string(runes)

		idxs := anagrams[normalized]
		idxs = append(idxs, i)
		anagrams[normalized] = idxs
	}

	result := make([][]string, 0, 10)

	for _, idxs := range anagrams {
		vals := make([]string, len(idxs))

		for i, idx := range idxs {
			vals[i] = strs[idx]
		}

		result = append(result, vals)
	}

	return result
}

func main() {

	var a ByRune = []rune("acb")

	sort.Sort(a)

	fmt.Printf("a: %v\n", string(a))

	input := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	ans := groupAnagrams(input)

	fmt.Printf("ans: %v\n", ans)

}
