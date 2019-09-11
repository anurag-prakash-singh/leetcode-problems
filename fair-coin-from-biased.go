package main

import (
	"fmt"
	"math/rand"
)

func biasedCoin() int {
	n := rand.Intn(100)

	if n < 60 {
		return 0
	}

	return 1
}

func fairCoin() int {
	for {
		coin1 := biasedCoin()
		coin2 := biasedCoin()

		if coin1 != coin2 {
			return coin1
		}
	}
}

type CoinFn func() int

func flip() {
	zeros := 0
	ones := 0

	var flipCoin CoinFn

	// flipCoin = biasedCoin
	flipCoin = fairCoin

	for i := 0; i < 100; i++ {
		flipResult := flipCoin()

		if flipResult == 0 {
			zeros++
		} else {
			ones++
		}
	}

	fmt.Printf("zeros: %d; ones: %d\n", zeros, ones)
}

func main() {
	flip()
}
