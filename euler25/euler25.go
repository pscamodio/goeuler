package main

import (
	"fmt"
	"math/big"
)

func main() {
	v1 := big.NewInt(1)
	v2 := big.NewInt(1)
	sum := big.NewInt(1)
	idx := 2
	for len(sum.String()) < 1000 {
		idx++
		sum.Add(v1, v2)
		v1.Set(v2)
		v2.Set(sum)
	}
	fmt.Println(idx)
}
