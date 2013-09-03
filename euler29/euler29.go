package main

/*
Consider all integer combinations of ab for 2 ≤ a ≤ 5 and 2 ≤ b ≤ 5:

    22=4, 23=8, 24=16, 25=32
    32=9, 33=27, 34=81, 35=243
    42=16, 43=64, 44=256, 45=1024
    52=25, 53=125, 54=625, 55=3125

If they are then placed in numerical order, with any repeats removed, we get the following sequence of 15 distinct terms:

4, 8, 9, 16, 25, 27, 32, 64, 81, 125, 243, 256, 625, 1024, 3125

How many distinct terms are in the sequence generated by ab for 2 ≤ a ≤ 100 and 2 ≤ b ≤ 100?
*/

import (
	"fmt"
	"math/big"
)

type intSlice []*big.Int

func pow(a, b int) *big.Int {
	tot := big.NewInt(1)
	base := big.NewInt(int64(a))
	for i := 0; i < b; i++ {
		tot = tot.Mul(tot, base)
	}
	return tot
}

func (m intSlice) contain(val *big.Int) bool {
	for v := range m {
		if m[v].Cmp(val) == 0 {
			return true
		}
	}
	return false
}

func main() {
	m := make(intSlice, 0, 9801)
	for a := 2; a <= 100; a++ {
		for b := 2; b <= 100; b++ {
			p := pow(a, b)
			if !m.contain(p) {
				m = append(m, p)
			}
		}
	}
	fmt.Println(len(m))
}