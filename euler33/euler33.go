package main

/*
The fraction 49/98 is a curious fraction, as an inexperienced mathematician in attempting to simplify it may incorrectly believe that 49/98 = 4/8, which is correct, is obtained by cancelling the 9s.

We shall consider fractions like, 30/50 = 3/5, to be trivial examples.

There are exactly four non-trivial examples of this type of fraction, less than one in value, and containing two digits in the numerator and denominator.

If the product of these four fractions is given in its lowest common terms, find the value of the denominator.
*/

import (
	"fmt"
)

func GCD(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

type fraction struct {
	num   int
	denum int
	div   int
	rest  int
	val   float64
	mod10 bool
}

func NewFraction(num, denum int) *fraction {
	d := fraction{num, denum, int(num / denum), num % denum, float64(num) / float64(denum), false}
	if num%10 == 0 && denum%10 == 0 {
		d.mod10 = true
	}
	return &d
}

func GetDigits(num int) []int {
	digits := make([]int, 2)
	digits[0] = num / 10
	digits[1] = num % 10
	return digits
}

func (fr *fraction) Simplify() (*fraction, bool) {
	numDigits := GetDigits(fr.num)
	denumDigits := GetDigits(fr.denum)
	var a, b int
	if numDigits[0] == denumDigits[0] {
		a = numDigits[1]
		b = denumDigits[1]
	} else if numDigits[0] == denumDigits[1] {
		a = numDigits[1]
		b = denumDigits[0]
	} else if numDigits[1] == denumDigits[0] {
		a = numDigits[0]
		b = denumDigits[1]
	} else if numDigits[1] == denumDigits[1] {
		a = numDigits[0]
		b = denumDigits[0]
	} else {
		return nil, false
	}
	if b == 0 {
		return nil, false
	}
	return NewFraction(a, b), true
}

func main() {
	numProd := 1
	denumProd := 1
	for num := 10; num <= 99; num++ {
		for denum := num + 1; denum <= 99; denum++ {
			d := NewFraction(num, denum)
			if d.mod10 {
				continue
			}
			simpl, valid := d.Simplify()
			if !valid {
				continue
			}
			if d.val < 1 && d.val == simpl.val {
				numProd *= simpl.num
				denumProd *= simpl.denum
			}
		}
	}
	gcd := GCD(numProd, denumProd)
	fmt.Println(numProd/gcd, "/", denumProd/gcd)
}
