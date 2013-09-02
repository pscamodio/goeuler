package main

import "fmt"
import "math/big"
import "strconv"

func main() {
	val := big.NewInt(1)
	val.Exp(big.NewInt(2), big.NewInt(1000), nil)
	s := val.String()
	tot := int64(0)
	for _, car := range s {
		num, err := strconv.ParseInt(string(car), 10, 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		tot += num
	}
	fmt.Println(tot)
}
