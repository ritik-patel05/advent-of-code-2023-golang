package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var cnts map[int]int64

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	cnts = map[int]int64{}

	currentId := 1

	for scanner.Scan() {
		line := scanner.Text()
		cnts[currentId]++
		getCardValue(currentId, line)
		currentId++
	}

	var ans int64
	for _, value := range cnts {
		ans += value
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println(ans)
}

func getCardValue(currentId int, card string) {
	words := strings.Fields(card)

	seenCnt := map[int]int64{}
	for _, word := range words {
		integer, err := strconv.Atoi(word)
		if err != nil {
			continue
		}

		seenCnt[integer]++
	}

	value := 0

	for _, total := range seenCnt {
		if total > 0 && total%2 == 0 {
			value++
		}
	}

	totalCurrentId := cnts[currentId]
	i := currentId + 1
	for i <= currentId+value {
		cnts[i] += totalCurrentId
		i++
	}
}
