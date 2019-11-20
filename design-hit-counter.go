package main

import (
	"fmt"
)

const (
	MaxNumCounts = 300
)

type HitCounter struct {
	head   int
	vHead  int
	tail   int
	vTail  int
	counts []int
}

/** Initialize your data structure here. */
func Constructor() HitCounter {
	return HitCounter{
		head:   MaxNumCounts - 1,
		vHead:  MaxNumCounts,
		tail:   0,
		vTail:  1,
		counts: make([]int, MaxNumCounts),
	}
}

func (this *HitCounter) resetCounts() {
	for i, _ := range this.counts {
		this.counts[i] = 0
	}
}

func (this *HitCounter) sum() int {
	s := 0

	for _, c := range this.counts {
		s += c
	}

	return s
}

/** Record a hit.
  @param timestamp - The current timestamp (in seconds granularity). */
func (this *HitCounter) Hit(timestamp int) {
	if timestamp-this.vHead >= MaxNumCounts {
		this.resetCounts()
		this.head = MaxNumCounts - 1
		this.tail = 0
		this.vHead = timestamp
		this.vTail = this.vHead - MaxNumCounts + 1
		this.counts[this.head] += 1

		return
	}

	if this.vHead < timestamp {
		for this.vHead < timestamp {
			this.vHead++
			this.vTail++

			this.head = (this.head + 1) % MaxNumCounts
			this.tail = (this.tail + 1) % MaxNumCounts

			this.counts[this.head] = 0
		}

		this.counts[this.head] = 1

		return
	}

	// this.vHead >= timestamp
	offset2 := (this.tail + (timestamp - this.vTail)) % MaxNumCounts
	this.counts[offset2] += 1
}

/** Return the number of hits in the past 5 minutes.
@param timestamp - The current timestamp (in seconds granularity).
*/
func (this *HitCounter) GetHits(timestamp int) int {
	if timestamp-this.vHead >= MaxNumCounts {
		this.resetCounts()
		this.head = MaxNumCounts - 1
		this.tail = 0
		this.vHead = timestamp
		this.vTail = this.vHead - MaxNumCounts + 1

		return 0
	}

	if this.vHead < timestamp {
		for this.vHead < timestamp {
			this.vHead++
			this.vTail++

			this.head = (this.head + 1) % MaxNumCounts
			this.tail = (this.tail + 1) % MaxNumCounts

			this.counts[this.head] = 0
		}

		this.counts[this.head] = 0

		return this.sum()
	}

	// this.vHead >= timestamp
	return this.sum()
}

/**
 * Your HitCounter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Hit(timestamp);
 * param_2 := obj.GetHits(timestamp);
 */

func test1() {
	hc := Constructor()

	hc.Hit(1)
	hc.Hit(2)
	hc.Hit(3)
	fmt.Printf("getHits(4) = %d (expected: 3)\n", hc.GetHits(4))

	hc.Hit(300)
	fmt.Printf("getHits(300) = %d (expected: 4)\n", hc.GetHits(300))

	fmt.Printf("getHits(301) = %d (expected: 3)\n", hc.GetHits(301))
}

func test2() {
	hc := Constructor()

	hc.Hit(1)
	hc.Hit(1)
	hc.Hit(1)
	hc.Hit(300)
	fmt.Printf("getHits(300) = %d (expected: 4)\n", hc.GetHits(300))

	hc.Hit(300)
	fmt.Printf("getHits(300) = %d (expected: 5)\n", hc.GetHits(300))

	hc.Hit(301)
	fmt.Printf("getHits(301) = %d (expected: 3)\n", hc.GetHits(301))
}

func main() {
	test1()
	// test2()
}
