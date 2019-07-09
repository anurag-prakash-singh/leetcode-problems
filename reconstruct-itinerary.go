package main

// reference: https://en.wikipedia.org/wiki/Eulerian_path

import (
	"fmt"
	"sort"
)

type AirportCodes []string

func (ac AirportCodes) Len() int {
	return len(ac)
}

func (ac AirportCodes) Swap(i, j int) {
	ac[i], ac[j] = ac[j], ac[i]
}

func (ac AirportCodes) Less(i, j int) bool {
	return ac[i] < ac[j] // strings.Compare(ac[i], ac[j]) < 0
}

func dfs(adjMat [][]int, from int, acs []string, traversed []string) []string {
	for to := 0; to < len(adjMat[from]); {
		if adjMat[from][to] <= 0 {
			to++
			continue
		}

		adjMat[from][to]--

		if adjMat[from][to] == 0 {
			adjMat[from][to] = -1
		}

		traversed = dfs(adjMat, to, acs, traversed)
	}

	return append(traversed, acs[from])
}

func reverse(strings []string) {
	for i, j := 0, len(strings)-1; i < j; i, j = i+1, j-1 {
		strings[i], strings[j] = strings[j], strings[i]
	}
}

func findItinerary(tickets [][]string) []string {
	acOffsets := make(map[string]int)

	for _, srcDest := range tickets {
		src, dest := srcDest[0], srcDest[1]

		acOffsets[src] = -1
		acOffsets[dest] = -1
	}

	acs := make([]string, len(acOffsets))
	acsOffset := 0

	for k, _ := range acOffsets {
		acs[acsOffset] = k
		acsOffset++
	}

	sort.Sort(AirportCodes(acs))

	for i, ac := range acs {
		acOffsets[ac] = i
	}

	fmt.Printf("acOffsets: %v\n", acOffsets)

	numAirports := len(acs)
	adjMat := make([][]int, numAirports)

	for i := 0; i < numAirports; i++ {
		adjMat[i] = make([]int, numAirports)
	}

	for _, ticket := range tickets {
		from, to := acOffsets[ticket[0]], acOffsets[ticket[1]]
		adjMat[from][to]++
	}

	traversed := dfs(adjMat, acOffsets["JFK"], acs, []string{})

	reverse(traversed)

	return traversed
}

func test1() {
	tickets := [][]string{
		[]string{"MUC", "LHR"},
		[]string{"JFK", "MUC"},
		[]string{"SFO", "SJC"},
		[]string{"LHR", "SFO"},
	}

	fmt.Printf("itinerary: %v\n", findItinerary(tickets))
}

func test2() {
	tickets := [][]string{
		[]string{"JFK", "LHR"},
		[]string{"LHR", "SFO"},
		[]string{"SFO", "JFK"},
		[]string{"JFK", "ATL"},
	}

	fmt.Printf("itinerary: %v\n", findItinerary(tickets))
}

func test3() {
	tickets := [][]string{
		[]string{"JFK", "AAA"},
		[]string{"AAA", "JFK"},
		[]string{"JFK", "BBB"},
		[]string{"JFK", "CCC"},
		[]string{"CCC", "JFK"},
	}

	fmt.Printf("itinerary: %v\n", findItinerary(tickets))
}

func test4() {
	tickets := [][]string{
		[]string{"JFK", "SFO"}, []string{"JFK", "ATL"}, []string{"SFO", "ATL"}, []string{"ATL", "JFK"}, []string{"ATL", "SFO"},
	}

	fmt.Printf("itinerary: %v\n", findItinerary(tickets))
}

func tests() {
	airportCodes := AirportCodes([]string{"JFK", "ATL", "SFO"})

	sort.Sort(airportCodes)

	fmt.Printf("sorted codes: %v\n", airportCodes)
}

func main() {
	// tests()
	// test1()
	// test2()
	// test3()
	test4()
}
