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
	"strings"
)

const (
	top int = 1000000
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

func cycleNext(num, max int) int {
	num++
	if num > max {
		num = 0
	}
	return num
}

func rotations(num int) []int {
	n_str := strconv.FormatInt(int64(num), 10)
	max := len(n_str) - 1
	rots := make([]int, max)
	for idx := 0; idx <= max-1; idx++ {
		str_slc := make([]string, max+1)
		for i, val := range n_str {
			str_slc[cycleNext(i, max)] = string(val)
		}
		new_n_str := strings.Join(str_slc, "")
		val, _ := strconv.ParseInt(new_n_str, 10, 0)
		if num == int(val) {
			return nil
		}
		rots[idx] = int(val)
		n_str = new_n_str
	}
	return rots
}

func contains(arr []int, val int) bool {
	idx := sort.SearchInts(arr, val)
	return idx < len(arr) && arr[idx] == val
}

func main() {
	primes := sieve(1000000)
	goods := make([]int, 0)
	for _, prime := range primes {
		if contains(goods, prime) {
			continue
		}
		rots := rotations(prime)
		good := true
		for _, rot := range rots {
			if !contains(primes, rot) {
				good = false
				break
			}
		}
		if good {
			goods = append(goods, prime)
			for _, rot := range rots {
				goods = append(goods, rot)
			}
		}
		sort.Ints(goods)
	}
	fmt.Println(goods)
	fmt.Println("sol ->", len(goods))
}
