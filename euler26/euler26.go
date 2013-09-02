package main

import (
	"fmt"
)

const (
	maxd int = 1000
)

func decimalCycleLengh(d int) int {
	rem := 1 % d
	idx := 0
	foundAtCycle := make(map[int]int)
	for {
		if rem == 0 {
			return 0
		}
		rem = rem * 10 % d
		line, ok := foundAtCycle[rem]
		if ok {
			return idx - line
		} else {
			foundAtCycle[rem] = idx
		}
		idx++
	}
}

func main() {
	longestCycle := 0
	numWithLongestCycle := 0
	for i := maxd; i >= 0; i-- {
		length := decimalCycleLengh(i)
		if length > longestCycle {
			longestCycle = length
			numWithLongestCycle = i
		}
		if longestCycle > i {
			break
		}
	}
	fmt.Println("Num With Longest Cycle: ", numWithLongestCycle, " Cycle: ", longestCycle)
}
