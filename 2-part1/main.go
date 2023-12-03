package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	TOTAL_RED_CUBES   = 12
	TOTAL_GREEN_CUBES = 13
	TOTAL_BLUE_CUBES  = 14
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var gameSum int
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

func isGamePossible(game string) (number int) {
	words := strings.Fields(game)

	gameString := words[1]
	gameString = gameString[:len(gameString)-1]
	gameNumber, err := strconv.Atoi(gameString)
	if err != nil {
		panic(err)
	}
	words = words[2:]

	var previousNumber int

	var (
		red   int
		green int
		blue  int
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
			if red > TOTAL_RED_CUBES || blue > TOTAL_BLUE_CUBES || green > TOTAL_GREEN_CUBES {
				return
			}
			red = 0
			green = 0
			blue = 0
		}

	}
	if red > TOTAL_RED_CUBES || blue > TOTAL_BLUE_CUBES || green > TOTAL_GREEN_CUBES {
		return
	}

	number = gameNumber
	return
}
