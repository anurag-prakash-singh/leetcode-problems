package main

import (
	"fmt"
	"strings"
)

type edge struct {
	from     string
	to       string
	dirLabel string // read as <from> is <dirLabel> of <to>
}

type vertexCoords struct {
	x int
	y int
}

type vertexAdjs struct {
	vertex    string
	fromEdges []edge
	toEdges   []edge
}

func processDirections(directions []string) []edge {
	edges := []edge{}

	for _, direction := range directions {
		dirParts := strings.Split(direction, " ")
		from, dirLabel, to := dirParts[0], dirParts[1], dirParts[2]
		edges = append(edges, edge{from: from, to: to, dirLabel: dirLabel})
	}

	return edges
}

func buildVertextAdjs(edges []edge) map(string, []vertexAdjs) {
	adjs := make(map(string, vertexAdjs))

	for _, e := range edges {
		f, t := e.from, e.to

		_, okf := adjs[f]
		_, okt := adjs[t]

		if !okf {
			adjs[f] = vertexAdjs{vertex: f, fromEdges: []edge{}, toEdges: []edge{}}
		}

		if !okt {
			adjs[t] = vertexAdjs{vertex: f, fromEdges: []edge{}, toEdges: []edge{}}
		}


	}
}

type testcase struct {
	directions []string
	expectedValidity bool
}

func test() {
	tcs := []testcase{
		testcase {
			directions: []string{"A NE B", "A N B"},
			expectedValidity: true,
		},
	}

	for _, tc := range tcs {
		edges := processDirections(tc.directions)

		fmt.Printf("edges: %v\n", edges)
	}
}

func main() {
	test()
}
