package main

import (
	"fmt"
)

const (
	maxToTest     int = 28123
	firstAbundant int = 12
)

func isAbundant(num int) bool {
	max := num/2 + 1
	tot := 1
	for i := 2; i < max; i++ {
		if num%i == 0 {
			tot += i
		}
	}
	return tot > num
}

func main() {
	abundants := make([]int, 0)
	for i := firstAbundant; i < maxToTest-firstAbundant; i++ {
		if isAbundant(i) {
			abundants = append(abundants, i)
		}
	}
	toTrash := make(map[int]bool)
	for a := range abundants {
	inner:
		for b := a; b < len(abundants); b++ {
			val := abundants[a] + abundants[b]
			if val < maxToTest {
				toTrash[val] = true
			} else {
				break inner
			}
		}
	}
	tot := 0
	for idx := 1; idx < maxToTest; idx++ {
		_, trash := toTrash[idx]
		if !trash {
			tot += idx
		}
	}
	fmt.Println(tot)
}
