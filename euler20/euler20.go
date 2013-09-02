package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	b := big.NewInt(1)
	for i := 1; i <= 100; i++ {
		b.Mul(b, big.NewInt(int64(i)))
	}
	s := b.String()
	tot := 0
	for _, car := range s {
		num, err := strconv.ParseInt(string(car), 10, 0)
		if err != nil {
			fmt.Println(err)
			return
		}
		tot += int(num)
	}
	fmt.Println(tot)
}
