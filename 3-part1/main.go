package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

	var grid [][]string

	for scanner.Scan() {
		line := scanner.Text()
		row := strings.Split(line, "")
		grid = append(grid, row)
	}

	sum := getPartSum(grid)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println(sum)
}

func getPartSum(grid [][]string) (sum int64) {
	for i := 0; i < len(grid); i++ {
		j := 0
		for j < len(grid[i]) {
			start := j
			for j < len(grid[i]) && unicode.IsDigit(rune(grid[i][j][0])) {
				j++
			}
			end := j - 1

			if end < start {
				j++
				continue
			}

			isSymbolPresent := false
			for k := start; k <= end; k++ {
				if hasSymbol(grid, i, k) {
					isSymbolPresent = true
					break
				}
			}

			if isSymbolPresent {
				number := ""
				for k := start; k <= end; k++ {
					number = number + string(grid[i][k][0])
				}
				numberInt, err := strconv.ParseInt(number, 10, 64)
				if err != nil {
					panic(err)
				}
				sum += numberInt
			}
		}
	}

	return
}

func hasSymbol(grid [][]string, i int, j int) bool {
	dirx := []int{1, 0, -1, 0, 1, 1, -1, -1}
	diry := []int{0, 1, 0, -1, 1, -1, 1, -1}

	for k := 0; k < 8; k++ {
		x := i + dirx[k]
		y := j + diry[k]
		if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) {
			char := grid[x][y][0]
			if char == '.' {
				continue
			} else if unicode.IsDigit(rune(char)) {
				continue
			} else {
				return true
			}
		}
	}

	return false
}
