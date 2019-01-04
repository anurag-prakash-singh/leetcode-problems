package main

import (
	"errors"
	"fmt"
	"math"
	"sort"
)

type Diagonal struct {
	x1, y1, x2, y2 float64
	length         float64
}

type ByLen []Diagonal

func (ds ByLen) Len() int {
	return len(ds)
}

func (ds ByLen) Swap(i, j int) {
	ds[i], ds[j] = ds[j], ds[i]
}

func (ds ByLen) Less(i, j int) bool {
	return ds[i].length < ds[j].length
}

func distance(x1, y1, x2, y2 float64) float64 {
	d := (x1-x2)*(x1-x2) + (y1-y2)*(y1-y2)
	return math.Sqrt(d)
}

func between(a, b, c float64) bool {
	if (c >= a && c <= b) || (c >= b && c <= a) {
		return true
	}

	return false
}

func intersection(d1, d2 Diagonal) (x float64, y float64, err error) {
	x1, y1 := d1.x1, d1.y1
	x2, y2 := d1.x2, d1.y2
	x3, y3 := d2.x1, d2.y1
	x4, y4 := d2.x2, d2.y2

	den := (x1-x2)*(y3-y4) - (y1-y2)*(x3-x4)

	if den == 0 {
		return 0, 0, errors.New("No intersection")
	}

	ix := ((x1*y2-y1*x2)*(x3-x4) - (x1-x2)*(x3*y4-y3*x4)) / den
	iy := ((x1*y2-y1*x2)*(y3-y4) - (y1-y2)*(x3*y4-y3*x4)) / den

	if between(x1, x2, ix) && between(x3, x4, ix) && between(y1, y2, iy) && between(y3, y4, iy) {
		return ix, iy, nil
	}

	return 0, 0, errors.New("No intersection")
}

func epseq(v1, v2 float64) bool {
	return (math.Abs(v1-v2) <= 0.00001)
}

func diagonalMidPoint(d Diagonal, x, y float64) bool {
	dist1 := distance(d.x1, d.y1, x, y)
	dist2 := distance(d.x2, d.y2, x, y)

	return epseq(dist1, dist2)
}

func minAreaFreeRect(points [][]int) float64 {
	n := len(points)

	diagonals := make([]Diagonal, (n*(n-1))/2)

	// fmt.Printf("num diags: %d\n", len(diagonals))

	k := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			d := Diagonal{x1: float64(points[i][0]), y1: float64(points[i][1]), x2: float64(points[j][0]), y2: float64(points[j][1])}
			d.length = distance(d.x1, d.y1, d.x2, d.y2)
			diagonals[k] = d

			k++
		}
	}

	byLen := ByLen(diagonals)
	sort.Sort(byLen)

	// fmt.Printf("diagonals: %v\n", diagonals)

	minArea := math.MaxFloat64
	rectFound := false

	for i := 0; i < byLen.Len(); i++ {
		for j := i + 1; j < byLen.Len(); j++ {
			if epseq(diagonals[i].length, diagonals[j].length) {
				// get intersection
				if ix, iy, err := intersection(diagonals[i], diagonals[j]); err == nil {
					if diagonalMidPoint(diagonals[i], ix, iy) && diagonalMidPoint(diagonals[j], ix, iy) {
						// This is a rectangle calculate area
						rectFound = true

						l := distance(diagonals[i].x1, diagonals[i].y1, diagonals[j].x1, diagonals[j].y1)
						w := distance(diagonals[i].x1, diagonals[i].y1, diagonals[j].x2, diagonals[j].y2)

						minArea = math.Min(minArea, l*w)
					}
				}
			} else {
				break
			}
		}
	}

	if !rectFound {
		return 0.0
	}

	return minArea
}

func test1() {
	points := [][]int{
		[]int{0, 1},
		[]int{2, 1},
		[]int{1, 1},
		[]int{1, 0},
		[]int{2, 0},
	}

	fmt.Printf("ans: %f\n", minAreaFreeRect(points))
}

func test2() {
	points := [][]int{
		[]int{0, 3},
		[]int{1, 2},
		[]int{3, 1},
		[]int{1, 3},
		[]int{2, 1},
	}

	fmt.Printf("ans: %f\n", minAreaFreeRect(points))
}

func test3() {
	points := [][]int{
		[]int{3, 1},
		[]int{1, 1},
		[]int{0, 1},
		[]int{2, 1},
		[]int{3, 3},
		[]int{3, 2},
		[]int{0, 2},
		[]int{2, 3},
	}

	fmt.Printf("ans: %f\n", minAreaFreeRect(points))
}

func test4() {
	points := [][]int{
		[]int{1, 2},
		[]int{2, 1},
		[]int{1, 0},
		[]int{0, 1},
	}

	fmt.Printf("ans: %f\n", minAreaFreeRect(points))
}

func main() {
	test1()
	test2()
	test3()
	test4()
}
