package main

import (
	"strconv"
	"unicode"
)

func adding(s string) int64 {
	var arrOfNumbers []int
	currentStr := ""

	for _, sym := range s {
		if unicode.IsDigit(sym) {
			currentStr = currentStr + string(sym)
		} else {
			number, _ := strconv.Atoi(currentStr)
			arrOfNumbers = append(arrOfNumbers, number)
			currentStr = ""
		}
	}

	sum := 0

	for i := range len(arrOfNumbers) {
		sum += arrOfNumbers[i]
	}

	return int64(sum)
}
