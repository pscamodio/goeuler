package main

import "fmt"

func IntToEnglish(val int) string {
	convMap := map[int]string{
		0:  "Zero",
		1:  "One",
		2:  "Two",
		3:  "Three",
		4:  "Four",
		5:  "Five",
		6:  "Six",
		7:  "Seven",
		8:  "Eight",
		9:  "Nine",
		10: "Ten",
		11: "Eleven",
		12: "Twelve",
		13: "Thirteen",
		14: "Fourteen",
		15: "Fifteen",
		16: "Sixteen",
		17: "Seventeen",
		18: "Eighteen",
		19: "Nineteen",
		20: "Twenty"}

	tenthMap := map[int]string{
		1: "Ten",
		2: "Twenty",
		3: "Thirty",
		4: "Forty",
		5: "Fifty",
		6: "Sixty",
		7: "Seventy",
		8: "Eighty",
		9: "Ninety"}

	hundred := " Hundred "
	thousand := " Thousand "
	and := "And "

	if val < 20 {
		return convMap[val]
	}
	if val < 100 {
		dec := val / 10
		rest := val % 10
		if rest != 0 {
			return tenthMap[dec] + "-" + IntToEnglish(val%10)
		} else {
			return tenthMap[dec]
		}
	}
	if val < 1000 {
		cent := val / 100
		rest := val % 100
		if rest != 0 {
			return convMap[cent] + hundred + and + IntToEnglish(rest)
		} else {
			return convMap[cent] + hundred
		}
	}
	if val < 100000 {
		t := val / 1000
		rest := val % 1000
		if rest != 0 {
			return convMap[t] + thousand + and + IntToEnglish(rest)
		} else {
			return convMap[t] + thousand
		}
	}

	return convMap[1]
}

func main() {
	tot := 0
	for x := 1; x <= 1000; x++ {
		inLetter := IntToEnglish(x)
		numDigits := 0
		for _, char := range inLetter {
			s := string(char)
			if s != "-" && s != " " {
				numDigits += 1
			}
		}
		tot += numDigits
	}
	fmt.Println(tot)
}
