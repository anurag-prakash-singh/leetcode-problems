package main

func trap(height []int) int {
	n := len(height)
	lo, hi := 0, n-1
	maxLeft, maxRight := 0, 0
	totalAccumulated := 0

	for lo < hi {
		if height[lo] < height[hi] {
			if height[lo] < maxLeft {
				totalAccumulated += maxLeft - height[lo]
			} else {
				maxLeft = height[lo]
			}

			lo++
		} else {
			if height[hi] < maxRight {
				totalAccumulated += maxRight - height[hi]
			} else {
				maxRight = height[hi]
			}

			hi--
		}
	}

	return totalAccumulated
}

func main() {

}
