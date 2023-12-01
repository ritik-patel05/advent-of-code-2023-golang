package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// Create a mapping of words to digits
var wordToDigit map[string]int

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	wordToDigit = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	var calibrationValue int64
	for scanner.Scan() {
		line := scanner.Text()
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
	for id, char := range line {
		if unicode.IsDigit(char) {
			calibrationValue = string(char)
			break
		}
		num, err := isWordNumberPrefix(id, line, false)
		if err == nil {
			calibrationValue = num
			break
		}
	}

	lastStringChar := ""
	for id, char := range line {
		if unicode.IsDigit(char) {
			lastStringChar = string(char)

		}
		num, err := isWordNumberPrefix(id, line, false)
		if err == nil {
			lastStringChar = num
		}
	}

	calibrationValue = fmt.Sprintf("%s%s", calibrationValue, string(lastStringChar))

	number, err := strconv.ParseInt(calibrationValue, 10, 64)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	return
}

func isWordNumberPrefix(index int, line string, getLast bool) (number string, err error) {

	for word, digit := range wordToDigit {
		// Check if the substring from the current index matches the word
		if strings.HasPrefix(line[index:], word) {
			if number == "" {
				number = strconv.Itoa(digit)
				if !getLast {
					return
				}
			} else {
				number = strconv.Itoa(digit)
			}
		}
	}

	err = fmt.Errorf("no digit found")
	return
}
