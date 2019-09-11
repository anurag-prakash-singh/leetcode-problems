package main

import "fmt"

type intSet struct {
	size     int
	elements map[int]struct{}
}

func NewIntSet() intSet {
	return intSet{size: 0, elements: make(map[int]struct{})}
}

func (s *intSet) Add(elem int) {
	if _, exists := s.elements[elem]; !exists {
		s.elements[elem] = struct{}{}
		s.size++
	}

	return
}

func (s *intSet) Remove(elem int) {
	if _, exists := s.elements[elem]; exists {
		delete(s.elements, elem)
		s.size--
	}

	return
}

func (s *intSet) Contains(elem int) bool {
	_, exists := s.elements[elem]

	return exists
}

func (s *intSet) GetAnyElement() (int, error) {
	if len(s.elements) == 0 {
		return 0, fmt.Errorf("set is empty")
	}

	var elem int
	for k, _ := range s.elements {
		elem = k

		break
	}

	return elem, nil
}

// Return true if cycle exists.
func checkCycleFrom(start int, adjs [][]int, whiteSet, graySet, blackSet intSet) bool {
	if blackSet.Contains(start) {
		return false
	}

	if graySet.Contains(start) {
		return true
	}

	if whiteSet.Contains(start) {
		whiteSet.Remove(start)
		graySet.Add(start)
	}

	// Explore all adjacencies
	for i := 0; i < len(adjs[start]); i++ {
		cycleExists := checkCycleFrom(adjs[start][i], adjs, whiteSet, graySet, blackSet)

		if cycleExists {
			return true
		}
	}

	graySet.Remove(start)
	blackSet.Add(start)

	return false
}

func checkCycle(adjs [][]int) bool {
	whiteSet, graySet, blackSet := NewIntSet(), NewIntSet(), NewIntSet()

	for i := 0; i < len(adjs); i++ {
		whiteSet.Add(i)
	}

	for {
		elem, err := whiteSet.GetAnyElement()

		if err != nil {
			break
		}

		if checkCycleFrom(elem, adjs, whiteSet, graySet, blackSet) {
			return true
		}
	}

	return false
}

func topologicalTraversal(start int, adjs [][]int, visited map[int]struct{}, traversal []int) []int {
	if _, exists := visited[start]; exists {
		return traversal
	}

	for _, to := range adjs[start] {
		traversal = topologicalTraversal(to, adjs, visited, traversal)
	}

	visited[start] = struct{}{}
	traversal = append(traversal, start)

	return traversal
}

func findOrder(numCourses int, prerequisites [][]int) []int {
	adjs := make([][]int, numCourses)

	for i, _ := range adjs {
		adjs[i] = []int{}
	}

	visited := make(map[int]struct{})
	leafCheck := make([]bool, numCourses)

	for i, _ := range leafCheck {
		leafCheck[i] = true
	}

	for _, pair := range prerequisites {
		from, to := pair[0], pair[1]

		adjs[from] = append(adjs[from], to)

		leafCheck[to] = false
	}

	if checkCycle(adjs) {
		return []int{}
	}

	order := []int{}

	for i := 0; i < numCourses; i++ {
		if leafCheck[i] {
			// leaf course - this is a course that no other course depends on
			order = topologicalTraversal(i, adjs, visited, order)
		}
	}

	for i := 0; i < numCourses; i++ {
		if _, exists := visited[i]; !exists {
			return []int{}
		}
	}

	return order
}

type testcase struct {
	numCourses int
	preReqs    [][]int
}

func tests() {
	testcases := []testcase{
		testcase{numCourses: 2, preReqs: [][]int{[]int{1, 0}}},
		testcase{numCourses: 4, preReqs: [][]int{[]int{1, 0}, []int{2, 0}, []int{3, 1}, []int{3, 2}}},
		testcase{numCourses: 3, preReqs: [][]int{[]int{0, 2}, []int{1, 2}, []int{2, 0}}},
	}

	for i, tc := range testcases {
		order := findOrder(tc.numCourses, tc.preReqs)

		fmt.Printf("Test case: %d; order: %v\n", i, order)
	}
}

func main() {
	// tests()
	s1 := NewIntSet()

	s1.Add(4)

	fmt.Printf("s1.size = %d\n", s1.size)
}
