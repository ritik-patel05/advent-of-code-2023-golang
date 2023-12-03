package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var gameSum int64
	for scanner.Scan() {
		game := scanner.Text()
		gameSum += isGamePossible(game)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println(gameSum)
}

func isGamePossible(game string) (number int64) {
	words := strings.Fields(game)

	words = words[2:]

	var (
		previousNumber int
		err            error
	)

	var (
		red   int
		green int
		blue  int
	)

	var (
		minRed   int
		minGreen int
		minBlue  int
	)

	for _, word := range words {
		lastChar := word[len(word)-1]
		if lastChar == ';' {
			word = word[:len(word)-1]
		} else if lastChar == ',' {
			word = word[:len(word)-1]
		}

		if word == "green" {
			green = previousNumber
		} else if word == "blue" {
			blue = previousNumber
		} else if word == "red" {
			red = previousNumber
		} else {
			previousNumber, err = strconv.Atoi(word)
			if err != nil {
				panic(err)
			}
		}

		if lastChar == ';' {
			if red > minRed {
				minRed = red
			}
			if blue > minBlue {
				minBlue = blue
			}
			if green > minGreen {
				minGreen = green
			}
			red = 0
			green = 0
			blue = 0
		}

	}

	if red > minRed {
		minRed = red
	}
	if blue > minBlue {
		minBlue = blue
	}
	if green > minGreen {
		minGreen = green
	}

	number = int64(minRed * minGreen * minBlue)
	return
}
