package main

/*
Surprisingly there are only three numbers that can be written as the sum of fourth powers of their digits:

    1634 = 14 + 64 + 34 + 44
    8208 = 84 + 24 + 04 + 84
    9474 = 94 + 44 + 74 + 44

As 1 = 14 is not a sum it is not included.

The sum of these numbers is 1634 + 8208 + 9474 = 19316.

Find the sum of all the numbers that can be written as the sum of fifth powers of their digits.
*/

import (
	"fmt"
)

const (
	lowLimit int64 = 4
	upLimit  int64 = 354294
)

func pow(base, exp int64) int64 {
	tot := int64(1)
	for i := int64(0); i < exp; i++ {
		tot *= base
	}
	return tot
}

func digits(num int64) []int64 {
	digits := make([]int64, 6, 6)
	div := int64(100000)
	for i := 0; i < 6; i++ {
		digits[i] = num / div
		num = num % div
		div = div / 10
	}
	return digits
}

func reduceDigits(digits []int64, exp int64) int64 {
	tot := int64(0)
	for _, digit := range digits {
		tot += pow(digit, exp)
	}
	return tot
}

func main() {
	tot := int64(0)
	for i := lowLimit; i <= upLimit; i++ {
		digits := digits(i)
		reduced := reduceDigits(digits, 5)
		if i == reduced {
			fmt.Println(i, digits, reduced)
			tot += i
		}
	}
	fmt.Println(tot)
}
