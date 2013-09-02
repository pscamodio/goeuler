package main

import (
	"fmt"
	"sort"
)

type LCode []int
type Seq []int

func nextLehemerCode(code LCode) bool {
	for idx := len(code) - 2; idx >= 0; idx-- {
		maxVal := len(code) - idx - 1
		if code[idx] < maxVal {
			code[idx]++
			return true
		} else {
			code[idx] = 0
		}
	}
	return false
}

func nthLegemerCode(nth, numDigits int) LCode {
	code := make(LCode, numDigits)
	for i := 1; i < nth; i++ {
		nextLehemerCode(code)
	}
	return code
}

func nthPermutation(nth int, digits_ Seq) Seq {
	digits := make(Seq, len(digits_))
	copy(digits, digits_)
	code := nthLegemerCode(nth, len(digits))
	s := make(Seq, len(digits))
	for idx, num := range code {
		sort.Ints(digits)
		s[idx] = digits[num]
		digits = append(digits[:num], digits[num+1:]...)
	}
	return s
}

func main() {
	seq := Seq{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	code := nthPermutation(1000000, seq)
	fmt.Println(code)
}
