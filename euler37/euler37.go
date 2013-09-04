package main

/*
The number 3797 has an interesting property. Being prime itself, it is possible to continuously remove digits from left to right, and remain prime at each stage: 3797, 797, 97, and 7. Similarly we can work from right to left: 3797, 379, 37, and 3.

Find the sum of the only eleven primes that are both truncatable from left to right and right to left.

NOTE: 2, 3, 5, and 7 are not considered to be truncatable primes.
*/

import (
	"fmt"
	"sort"
	"strconv"
)

func sieve(max int) []int {
	m_ok := make([]bool, max, max)
	l := make([]int, 0, max)
	for i := 2; i < max; i++ {
		m_ok[i] = true
	}
	for i := 2; i < max; i++ {
		if m_ok[i] == true {
			l = append(l, i)
			for j := 2; j*i < max; j++ {
				m_ok[i*j] = false
			}
		}
	}
	return l
}

func contains(arr []int, val int) bool {
	idx := sort.SearchInts(arr, val)
	return idx < len(arr) && arr[idx] == val
}

func testTrunks(val int, primes []int) bool {
	str := strconv.FormatInt(int64(val), 10)
	if len(str) <= 1 {
		return false
	}
	for i := 1; i < len(str); i++ {
		val1, _ := strconv.ParseInt(str[i:], 10, 0)
		val2, _ := strconv.ParseInt(str[:len(str)-i], 10, 0)

		if !contains(primes, int(val1)) || !contains(primes, int(val2)) {
			return false
		}
	}
	return true
}

func main() {
	primes := sieve(1000000)
	tot := 0
	for i := len(primes) - 1; i >= 0; i-- {
		if testTrunks(primes[i], primes) {
			tot += primes[i]
		}
	}
	fmt.Println(tot)
}
