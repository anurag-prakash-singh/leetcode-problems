package main

import (
	"fmt"
	"math"
)
  
func Evaporator(content float64, evapPerDay int, threshold int) int {
	// your code
	n := math.Log10(float64(threshold) / 100) / math.Log10(1 - float64(evapPerDay) / 100.0)
	
	return int(math.Ceil(n))
}

func main() {
	fmt.Printf("%d\n", Evaporator(10, 10, 10))
	fmt.Printf("%d\n", Evaporator(10, 10, 5))
	fmt.Printf("%d\n", Evaporator(100, 5, 5))
}
