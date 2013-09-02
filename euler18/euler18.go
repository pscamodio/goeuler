package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	data := `75
95 64
17 47 82
18 35 87 10
20 04 82 47 65
19 01 23 75 03 34
88 02 77 73 07 63 67
99 65 04 28 06 16 70 92
41 41 26 56 83 40 80 70 33
41 48 72 33 47 32 37 16 94 29
53 71 44 65 25 43 91 52 97 51 14
70 11 33 28 77 73 17 78 39 68 17 57
91 71 52 38 17 14 91 43 58 50 27 29 48
63 66 04 68 89 53 67 30 73 16 69 87 40 31
04 62 98 27 23 09 70 98 73 93 38 53 60 04 23`
	lines := strings.Split(data, "\n")
	tree := make([][]int, 0)
	for _, line := range lines {
		row := make([]int, 0)
		for _, num := range strings.Split(line, " ") {
			val, err := strconv.ParseInt(num, 10, 0)
			if err != nil {
				fmt.Println(err)
			}
			row = append(row, int(val))
		}
		tree = append(tree, row)
	}
	/*
		Reducing the tree, starting from the second row from bottom i sum
		the value of a number with the max of his leaf and then repeate going
		up row by row
	*/
	for r := len(tree) - 2; r >= 0; r-- {
		for idx := range tree[r] {
			if tree[r+1][idx] > tree[r+1][idx+1] {
				tree[r][idx] += tree[r+1][idx]
			} else {
				tree[r][idx] += tree[r+1][idx+1]
			}
		}
	}
	//Now the edge has the max value
	fmt.Println(tree[0][0])
}
