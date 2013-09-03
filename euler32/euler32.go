package main

/*
We shall say that an n-digit number is pandigital if it makes use of all the digits 1 to n exactly once; for example, the 5-digit number, 15234, is 1 through 5 pandigital.

The product 7254 is unusual, as the identity, 39 × 186 = 7254, containing multiplicand, multiplier, and product is 1 through 9 pandigital.

Find the sum of all products whose multiplicand/multiplier/product identity can be written as a 1 through 9 pandigital.
HINT: Some products can be obtained in more than one way so be sure to only include it once in your sum.
*/

import (
	"fmt"
	"sort"
	"strconv"
)

type data struct {
	moltiplicand int64
	multiplier   int64
	product      int64
}

func NewData(moltiplicand, multiplier int64) *data {
	n := data{moltiplicand, multiplier, moltiplicand * multiplier}
	return &n
}

func (d *data) getDigits() []int {
	str := strconv.FormatInt(d.moltiplicand, 10) + strconv.FormatInt(d.multiplier, 10) + strconv.FormatInt(d.product, 10)
	digits := make([]int, len(str))
	for idx, r := range str {
		num, _ := strconv.ParseInt(string(r), 10, 0)
		digits[idx] = int(num)
	}
	return digits
}

func isPandigital(digits []int) bool {
	if len(digits) != 9 {
		return false
	}
	sort.Ints(digits)
	for i := 1; i <= 9; i++ {
		if digits[i-1] != i {
			return false
		}
	}
	return true
}

func main() {
	tot := int64(0)
	used := make(map[int64]bool)
	for a := int64(1); a < 100; a++ {
		for b := a + 1; b < 98765432; b++ {
			d := NewData(a, b)
			digits := d.getDigits()
			if isPandigital(digits) {
				_, ok := used[d.product]
				if !ok {
					tot += d.product
					used[d.product] = true
				}
			}
			if len(digits) > 9 {
				break
			}
		}
	}
	fmt.Println(tot)
}
