package main

import "fmt"

func sumOfDivisors(num int) int {
	tot := 1
	for div := 2; div < num; div++ {
		if num%div == 0 {
			tot += div
		}
	}
	return tot
}

func main() {
	amicables := make(map[int]bool)
	for i := 1; i < 10000; i++ {
		_, ok := amicables[i]
		if ok {
			continue
		}
		sum := sumOfDivisors(i)
		revSum := sumOfDivisors(sum)
		if revSum == i && i != sum {
			amicables[i] = true
			amicables[sum] = true
		}
	}
	tot := 0
	for amicable, _ := range amicables {
		tot += amicable
	}
	fmt.Println(tot)
}
