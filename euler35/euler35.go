package main

/*
The number, 197, is called a circular prime because all rotations of the digits: 197, 971, and 719, are themselves prime.

There are thirteen such primes below 100: 2, 3, 5, 7, 11, 13, 17, 31, 37, 71, 73, 79, and 97.

How many circular primes are there below one million?
*/
import (
	"fmt"
	"sort"
	"strconv"
)

const (
	top int64 = 1000000
)

func sieve(max int64) []int64 {
	m_ok := make([]bool, max, max)
	l := make([]int64, 0, max)
	for i := int64(2); i < max; i++ {
		m_ok[i] = true
	}
	for i := int64(2); i < max; i++ {
		if m_ok[i] == true {
			l = append(l, i)
			for j := int64(2); j*i < max; j++ {
				m_ok[i*j] = false
			}
		}
	}
	return l
}

func rotations(num int64) []int64 {
	digits := strconv.FormatInt(num, 10)
	for idx := 1; idx < len(digits); i++{
		
	}
}

func main() {
	primes := sieve(top)
	for _,prime := primes {
	}
}
