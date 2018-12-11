package main

import "fmt"

type xn struct {
	x float64
	n int
}

func cachedMyPow(x float64, n int, cache map[xn]float64) float64 {
	if x == 1 {
		return 1
	}

	if n == 1 {
		return x
	}

	if n == 0 {
		return 1
	}

	if v, ok := cache[xn{x, n}]; ok {
		return v
	}

	switch n % 2 {
	case 0:
		cache[xn{x, n}] = cachedMyPow(x, n/2, cache) * cachedMyPow(x, n/2, cache)
	case 1:
		cache[xn{x, n}] = x * cachedMyPow(x, n/2, cache) * cachedMyPow(x, n/2, cache)
	}

	return cache[xn{x, n}]
}

func myPow(x float64, n int) float64 {
	cache := make(map[xn]float64)

	cache[xn{x, 1}] = x
	cache[xn{x, 0}] = 1

	if n < 0 {
		return 1 / cachedMyPow(x, -n, cache)
	}

	return cachedMyPow(x, n, cache)
}

func main() {
	fmt.Printf("2.1 ^ 3 = %f\n", myPow(2.1, 3))
	fmt.Printf("2 ^ -2 = %f\n", myPow(2, -2))
}
