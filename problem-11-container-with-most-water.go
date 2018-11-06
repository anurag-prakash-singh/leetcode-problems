package main

import (
	"fmt"
)

func maxArea(height []int) int {
	maxArea := 0
	i, j := 0, len(height)-1

	for i < j {
		h := height[i]

		if h > height[j] {
			h = height[j]
		}

		area := h * (j - i)

		if area > maxArea {
			maxArea = area
		}

		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}

	return maxArea
}

func main() {

	heights := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}

	fmt.Printf("maxArea: %d\n", maxArea(heights))

}
