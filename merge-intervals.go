package main

import (
	"fmt"
	"sort"
)

type Interval struct {
	Start int
	End   int
}

type ByStart []Interval

func (intvls ByStart) Len() int {
	return len(intvls)
}

func (intvls ByStart) Swap(i, j int) {
	intvls[i], intvls[j] = intvls[j], intvls[i]
}

func (intvls ByStart) Less(i, j int) bool {
	return intvls[i].Start < intvls[j].Start
}

func intervalsIntersect(intvl1, intvl2 Interval) bool {
	return ((intvl1.Start >= intvl2.Start && intvl1.Start <= intvl2.End) || (intvl1.End >= intvl2.Start && intvl1.End <= intvl2.End))
}

func max(a, b int) int {
	if a >= b {
		return a
	}

	return b
}

func merge(intervals []Interval) []Interval {
	sort.Sort(ByStart(intervals))

	mergedIntervals := make([]Interval, 0, len(intervals))

	if len(intervals) == 0 {
		return intervals
	}

	var currInterval Interval

	for i, interval := range intervals {
		if i == 0 {
			currInterval = interval

			continue
		}

		if intervalsIntersect(currInterval, interval) || intervalsIntersect(interval, currInterval) {
			currInterval.End = max(currInterval.End, interval.End)
		} else {
			mergedIntervals = append(mergedIntervals, currInterval)
			currInterval = interval
		}
	}

	mergedIntervals = append(mergedIntervals, currInterval)

	return mergedIntervals
}

func test1() {
	intvls := []Interval{
		Interval{1, 3},
		Interval{2, 6},
		Interval{8, 10},
		Interval{15, 18},
	}

	mergedIntervals := merge(intvls)

	fmt.Printf("merged: %v\n", mergedIntervals)
}

func test2() {
	intvls := []Interval{
		Interval{1, 4},
		Interval{4, 5},
	}

	mergedIntervals := merge(intvls)

	fmt.Printf("merged: %v\n", mergedIntervals)
}

func test3() {
	intvls := []Interval{
		Interval{1, 4},
	}

	mergedIntervals := merge(intvls)

	fmt.Printf("merged: %v\n", mergedIntervals)
}

func test4() {
	intvls := []Interval{}

	mergedIntervals := merge(intvls)

	fmt.Printf("merged: %v\n", mergedIntervals)
}

func test5() {
	intvls := []Interval{
		Interval{1, 4},
		Interval{2, 3},
	}

	mergedIntervals := merge(intvls)

	fmt.Printf("merged: %v\n", mergedIntervals)
}

func main() {

	// intvls := ByStart([]interval{interval{4, 3}, interval{5, 2}, interval{1, 10}})

	// sort.Sort(intvls)

	// fmt.Printf("sorted: %v\n", intvls)

	test1()
	test2()
	test3()
	test4()
	test5()
}
