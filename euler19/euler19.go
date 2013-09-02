package main

import (
	"fmt"
	"time"
)

func main() {
	loc, err := time.LoadLocation("")
	if err != nil {
		fmt.Println(err)
		return
	}
	count := 0
	for year := 1901; year <= 2000; year++ {
		for month := 1; month <= 12; month++ {
			day := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, loc)
			if day.Weekday() == time.Sunday {
				count++
			}
		}
	}
	fmt.Println(count)
}
