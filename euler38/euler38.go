package main

/*
Take the number 192 and multiply it by each of 1, 2, and 3:

    192 × 1 = 192
    192 × 2 = 384
    192 × 3 = 576

By concatenating each product we get the 1 to 9 pandigital, 192384576. We will call 192384576 the concatenated product of 192 and (1,2,3)

The same can be achieved by starting with 9 and multiplying by 1, 2, 3, 4, and 5, giving the pandigital, 918273645, which is the concatenated product of 9 and (1,2,3,4,5).

What is the largest 1 to 9 pandigital 9-digit number that can be formed as the concatenated product of an integer with (1,2, ... , n) where n > 1?
*/

import (
	"fmt"
	"runtime"
	"strconv"
)

const (
	startProd int = 2
	minBase   int = 2
	maxBase   int = 10000
)

func intToStr(num int) string {
	return strconv.FormatInt(int64(num), 10)
}

func strToInt(num string) int {
	val, _ := strconv.ParseInt(num, 10, 0)
	return int(val)
}

func concat(a string, b string) string {
	return a + b
}

func catenatedProductGenerator(base int, products chan int, end chan int) {
	prod := 2
	tot := intToStr(base)
	for {
		tot = concat(tot, intToStr(base*prod))
		if len(tot) == 9 {
			products <- strToInt(tot)
			end <- 1
			return
		} else if len(tot) > 9 {
			end <- 1
			return
		}
		prod++
	}
}

func isPandigital(num int) bool {
	str := intToStr(num)
	digits := make([]bool, 9, 9)
	for _, digit := range str {
		idx := strToInt(string(digit))
		if idx != 0 {
			digits[idx-1] = true
		}
	}
	for _, test := range digits {
		if !test {
			return false
		}
	}
	return true
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	maxRoutines := runtime.NumCPU()
	products := make(chan int)
	end := make(chan int)
	numRoutines := 0
	pandigital := 123456789
	i := minBase
	for {
		for i <= maxBase && numRoutines <= maxRoutines {
			i++
			numRoutines++
			go catenatedProductGenerator(i, products, end)
		}
		if i >= maxBase && numRoutines == 0 {
			break
		}
		select {
		case p := <-products:
			if isPandigital(p) && p > pandigital {
				pandigital = p
			}
		case <-end:
			numRoutines--
		}
	}
	fmt.Println(pandigital)
}
