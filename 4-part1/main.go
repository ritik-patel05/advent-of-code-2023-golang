package main

import (
	"bufio"
	"fmt"
	"os"
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

	var sum int64

	for scanner.Scan() {
		line := scanner.Text()
		sum += getCardValue(line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println(sum)
}

func getCardValue(card string) (value int64) {
	words := strings.Fields(card)

	seenCnt := map[string]int64{}
	for _, word := range words {
		seenCnt[word]++
	}

	for _, total := range seenCnt {
		if total > 0 && total%2 == 0 {
			if value == 0 {
				value = 1
			} else {
				value *= int64(2)
			}
		}
	}

	return
}
