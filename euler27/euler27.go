package main

/*
Euler discovered the remarkable quadratic formula:

n² + n + 41

It turns out that the formula will produce 40 primes for the consecutive values n = 0 to 39. However, when n = 40, 402 + 40 + 41 = 40(40 + 1) + 41 is divisible by 41, and certainly when n = 41, 41² + 41 + 41 is clearly divisible by 41.

The incredible formula  n² − 79n + 1601 was discovered, which produces 80 primes for the consecutive values n = 0 to 79. The product of the coefficients, −79 and 1601, is −126479.

Considering quadratics of the form:

    n² + an + b, where |a| < 1000 and |b| < 1000

    where |n| is the modulus/absolute value of n
    e.g. |11| = 11 and |−4| = 4

Find the product of the coefficients, a and b, for the quadratic expression that produces the maximum number of primes for consecutive values of n, starting with n = 0.
*/

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type data struct {
	A int64
	B int64
}

const (
	maxa int64 = 1000
	maxb int64 = 1000
)

func suppressWarnings() {
	_ = fmt.Print
	_ = ioutil.WriteFile
	_ = strings.Join
}

func isPrime(num int64) bool {
	if num <= 0 {
		return false
	}
	for i := int64(2); i <= int64(num/2); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func numQuadPrimes(num data) int64 {
	tot := int64(0)
	n := int64(0)
	for {
		prime := n*n + num.A*n + num.B
		if isPrime(prime) {
			tot++
		} else {
			break
		}
		n++
	}
	return tot
}

func main() {
	maxd := data{0, 0}
	maxPrimes := int64(0)
	for a := -maxa + 1; a < maxa; a++ {
		for b := -maxb + 1; b < maxb; b++ {
			d := data{a, b}
			newMax := numQuadPrimes(d)
			if newMax > maxPrimes {
				maxd = d
				maxPrimes = newMax
				fmt.Println("New best", maxd, "->", newMax)
			}
		}
	}
	fmt.Println("Better Data:", maxd)
	fmt.Println("Tot Primes:", maxPrimes)
	fmt.Println("Product:", maxd.A*maxd.B)
}
