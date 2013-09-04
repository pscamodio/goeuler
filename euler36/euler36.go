package main

/*
The decimal number, 585 = 10010010012 (binary), is palindromic in both bases.

Find the sum of all numbers, less than one million, which are palindromic in base 10 and base 2.

(Please note that the palindromic number, in either base, may not include leading zeros.)
*/

import (
	"fmt"
	"strconv"
)

const (
	max int = 1000000
)

func Reverse(s string) string {
	n := len(s)
	runes := make([]rune, n)
	for _, rune := range s {
		n--
		runes[n] = rune
	}
	return string(runes[n:])
}

func isDecPalindrome(num int) bool {
	return isPalindrome(strconv.FormatInt(int64(num), 10))
}

func isBinPalindrome(num int) bool {
	return isPalindrome(strconv.FormatInt(int64(num), 2))
}

func isPalindrome(s string) bool {
	return s == Reverse(s)
}

func main() {
	tot := 0
	for i := 1; i < max; i++ {
		if isDecPalindrome(i) && isBinPalindrome(i) {
			tot += i
		}
	}
	fmt.Println(tot)
}
