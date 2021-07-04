package day3

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func Move(input [][]int, fromX, fromY, right, down int) (nextX, nextY, value int) {
	nextY = (fromY + down)

	if nextY >= len(input) {
		return 0, 0, -1
	}

	columns := len(input[nextY])
	nextX = (fromX + right) % columns

	return nextX, nextY, input[nextY][nextX]
}

func CountTrees(input [][]int, right, down int) int {
	trees := 0
	fromX := 0
	fromY := 0
	value := 0

	for value >= 0 {
		trees += value
		fromX, fromY, value = Move(input, fromX, fromY, right, down)
	}

	return trees
}

func ParseInput(input io.Reader) [][]int {
	result := [][]int{}

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		result = append(result, parseLine(scanner.Text()))
	}

	return result
}

func parseLine(line string) []int {
	parsedLine := []int{}
	for _, value := range strings.Split(line, "") {
		parsedLine = append(parsedLine, parseValue(value))
	}
	return parsedLine
}

func parseValue(value string) int {
	if value == "#" {
		return 1
	}
	return 0
}

func SolvePart1(inputPath string) int {
	input, err := os.Open(inputPath)

	if err != nil {
		panic(err)
	}

	defer input.Close()

	return CountTrees(ParseInput(input), 3, 1)
}

func SolvePart2(inputPath string) int {
	input, err := os.Open(inputPath)

	if err != nil {
		panic(err)
	}

	defer input.Close()

	parsedInput := ParseInput(input)

	return CountTrees(parsedInput, 1, 1) *
		CountTrees(parsedInput, 3, 1) *
		CountTrees(parsedInput, 5, 1) *
		CountTrees(parsedInput, 7, 1) *
		CountTrees(parsedInput, 1, 2)
}
