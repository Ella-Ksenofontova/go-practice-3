package main

import (
	"math"
	"strconv"
	"strings"
)

func getResultOfDivision(csvString string) string {
	arrOfStrings := strings.Split(csvString, ";")
	firstNumberAsStr := strings.ReplaceAll(arrOfStrings[0], ",", ".")
	secondNumberAsStr := strings.ReplaceAll(arrOfStrings[1], ",", ".")

	firstNumberAsStr = strings.ReplaceAll(firstNumberAsStr, " ", "")
	secondNumberAsStr = strings.ReplaceAll(secondNumberAsStr, " ", "")

	firstNumber, _ := strconv.ParseFloat(firstNumberAsStr, 32)
	secondNumber, _ := strconv.ParseFloat(secondNumberAsStr, 32)

	result := math.Round(firstNumber/secondNumber*10_000) / 10_000
	resultAsStr := strconv.FormatFloat(result, 'f', 4, 32)

	return resultAsStr
}
