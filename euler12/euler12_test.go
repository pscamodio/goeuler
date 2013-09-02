package euler12

import "testing"
import "fmt"

func TestTriangleGenerator(t *testing.T) {
	//Test again first 10 numbers
	num := make(chan int)
	end := make(chan bool, 1)
	go TriangleNumGenerator(num, end)
	toTest := []int{1, 3, 6, 10, 15, 21, 28, 36, 45, 55}
	for i := 0; i < 10; i++ {
		val := <-num
		if val != toTest[i] {
			t.Fail()
		}
	}
	end <- true
}

func TestFactorCalculator(t *testing.T) {
	in := make(chan int)
	out := make(chan Result)
	go FactorsCalcolator(in, out)
	toTest := []int{1, 3, 6, 9, 12}
	for _, num := range toTest {
		in <- num
		factors := <-out
		fmt.Println(factors)
	}
}
