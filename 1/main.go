package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var calibrationValue int64
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		calibrationValue += processCalibrationLine(line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println(calibrationValue)
}

func processCalibrationLine(line string) (number int64) {
	calibrationValue := ""
	for _, char := range line {
		if unicode.IsDigit(char) {
			calibrationValue = string(char)
			break
		}
	}

	length := len(line)
	for i := length - 1; i >= 0; i-- {
		char := rune(line[i])
		if unicode.IsDigit(char) {
			calibrationValue = fmt.Sprintf("%s%s", calibrationValue, string(char))
			break
		}
	}

	number, err := strconv.ParseInt(calibrationValue, 10, 64)
	fmt.Println(number)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	return
}
