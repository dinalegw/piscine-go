package piscine

import (
	"fmt"
)

func DealAPackOfCards(deck []int) {
	// 4 players, each gets 3 cards (12 cards total)
	for i := 0; i < 4; i++ {
		fmt.Printf("Player %d: ", i+1)

		start := i * 3
		end := start + 3

		for j := start; j < end; j++ {
			if j == end-1 {
				fmt.Printf("%d", deck[j])
			} else {
				fmt.Printf("%d, ", deck[j])
			}
		}

		fmt.Printf("\n")
	}
}
