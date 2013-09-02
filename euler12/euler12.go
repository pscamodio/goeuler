package euler12

func TriangleNumGenerator(num chan int, end chan bool) {
	idx := 1
	n := 0
	looping := true
	for looping {
		n += idx
		idx += 1
		num <- n
		select {
		case <-end:
			looping = false
		default:
			looping = true
		}
	}
}

type Result struct {
	Num     int
	Factors []int
}

func FactorsCalcolator(input chan int, output chan Result) {
	for {
		num := <-input
		r := Result{num, nil}
		set := make(map[int]bool)
		for i := 1; i < num/2+1; i++ {
			if num%i == 0 {
				set[i] = true
				div := num / i
				set[div] = true
			}
		}
		for key, _ := range set {
			r.Factors = append(r.Factors, key)
		}
		output <- r
	}
}
