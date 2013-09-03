package main

/*
145 is a curious number, as 1! + 4! + 5! = 1 + 24 + 120 = 145.

Find the sum of all numbers which are equal to the sum of the factorial of their digits.

Note: as 1! = 1 and 2! = 2 are not sums they are not included.
*/

import (
	"fmt"
)

func suppressWarning() {
	_ = fmt.Print
}

const (
	max int64 = 50000
)

func digits(num int64) []int64 {
	digits := make([]int64, 0)
	div := int64(1000000)
	started := false
	for i := 0; i < 7; i++ {
		val := num / div
		num = num % div
		div = div / 10
		if val == 0 && !started {
			continue
		}
		started = true
		digits = append(digits, val)
	}
	return digits
}

func factorize(num int64) int64 {
	tot := int64(1)
	for i := int64(2); i <= num; i++ {
		tot *= i
	}
	return tot
}

func factorizeReduce(nums []int64) int64 {
	tot := int64(0)
	for _, num := range nums {
		tot += factorize(num)
	}
	return tot
}

func main() {
	tot := int64(0)
	for i := int64(10); i < max; i++ {
		if i == factorizeReduce(digits(i)) {
			tot += i
		}
	}
	fmt.Println(tot)
}
