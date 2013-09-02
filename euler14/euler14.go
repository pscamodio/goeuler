package main

import (
	"fmt"
	"runtime"
)

type Result struct {
	Val int
	L   int
}

func chainLengthCalc(input chan int, output chan Result) {
	for {
		val := <-input
		orig := val
		length := 0
		for val != 1 {
			length += 1
			if val%2 == 0 {
				val = val / 2
			} else {
				val = 3*val + 1
			}
		}
		output <- Result{orig, length}
	}
}

func main() {
	maxProcs := 1
	runtime.GOMAXPROCS(maxProcs)
	input := make(chan int)
	output := make(chan Result)
	for i := 0; i < maxProcs; i++ {
		go chainLengthCalc(input, output)
	}
	num := 1000000
	max := 0
	maxLength := 0
	for num > 2 {
		select {
		case l := <-output:
			if l.L > maxLength {
				max = l.Val
				maxLength = l.L
			}
		case input <- num:
			num -= 1
		}
	}
	fmt.Println(max, "-", maxLength)

}
