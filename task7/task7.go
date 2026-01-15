package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

func findSumFromFile(fileToReadFrom string) {
	file, _ := os.Open(fileToReadFrom)
	reader := bufio.NewReader(file)

	var numbers []int
	var errExternal error = nil

	for errExternal == nil {
		number, err := reader.ReadString('\n')
		if len(number) > 1 && number[len(number)-2] == '\r' {
			number = number[:len(number)-2]
		} else if len(number) > 0 && number[len(number)-1] == '\n' {
			number = number[:len(number)-1]
		}

		if err == nil || err == io.EOF {
			numberAsInt, _ := strconv.Atoi(number)
			numbers = append(numbers, numberAsInt)
		}

		errExternal = err
	}

	sum := 0

	for i := range len(numbers) {
		sum += numbers[i]
	}

	file.Close()

	resultFile, _ := os.Create("./result_task_7.txt")
	writer := bufio.NewWriter(resultFile)
	writer.WriteString(strconv.Itoa(sum))
	writer.Flush()

	resultFile.Close()
}

func main() {
	findSumFromFile("./numbers(task 7).txt")
}
