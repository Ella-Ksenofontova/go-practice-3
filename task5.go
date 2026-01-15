package main

import (
	"strconv"
)

func fn(n uint) uint {
	nAsString := strconv.Itoa(int(n))
	resultString := ""

	for _, sym := range nAsString {
		digit, _ := strconv.Atoi(string(sym))

		if digit%2 == 0 && digit != 0 {
			resultString += strconv.Itoa(digit)
		}
	}

	if len(resultString) == 0 {
		return 100
	}

	resultNumber, _ := strconv.Atoi(resultString)
	return uint(resultNumber)
}
