package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func AlphaValue(s string) int {

	charValue := map[rune]int{
		'A': 1,
		'B': 2,
		'C': 3,
		'D': 4,
		'E': 5,
		'F': 6,
		'G': 7,
		'H': 8,
		'I': 9,
		'J': 10,
		'K': 11,
		'L': 12,
		'M': 13,
		'N': 14,
		'O': 15,
		'P': 16,
		'Q': 17,
		'R': 18,
		'S': 19,
		'T': 20,
		'U': 21,
		'V': 22,
		'W': 23,
		'X': 24,
		'Y': 25,
		'Z': 26,
	}

	tot := 0
	for _, r := range s {
		tot += charValue[r]
	}
	return tot
}

func main() {
	data, err := ioutil.ReadFile("names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	names := strings.Split(string(data), ",")
	for idx := range names {
		names[idx] = strings.Trim(names[idx], "\"")
	}
	sort.Strings(names)
	tot := 0
	for idx, name := range names {
		tot += (idx + 1) * AlphaValue(name)
	}
	fmt.Println(tot)
}
