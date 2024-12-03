package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"unicode"
)

func parseInput() (string, error) {
	file, err := os.Open("./input.txt")
	if err != nil {
		return "", err
	}

	bytes, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func getOperands(input string) ([2]int, error) {
	if len(input) <= 4 {
		return [2]int{}, errors.New("too short")
	}

	if input[0] != 'm' || input[1] != 'u' || input[2] != 'l' || input[3] != '(' {
		return [2]int{}, errors.New("not a mul")
	}

	isFirst := true
	firstStr := ""
	secondStr := ""
	for _, char := range input[4:] {
		if unicode.IsDigit(char) {
			if isFirst {
				firstStr += string(char)
			} else {
				secondStr += string(char)
			}
		} else if char == ',' {
			isFirst = false
		} else if char == ')' {
			first, err := strconv.Atoi(firstStr)
			if err != nil {
				return [2]int{}, err
			}
			second, err := strconv.Atoi(secondStr)
			if err != nil {
				return [2]int{}, err
			}
			return [2]int{first, second}, nil
		} else {
			return [2]int{}, errors.New("unknown character: " + string(char))
		}
	}
	return [2]int{}, errors.New("unexpected end of input")
}

func solvePart1() {
	input, err := parseInput()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var operands [][2]int
	for index := range input {
		ops, err := getOperands(input[index:])
		if err != nil {
			continue
		}
		operands = append(operands, ops)
	}

	var res int64 = 0
	for _, ops := range operands {
		res += int64(ops[0] * ops[1])
	}

	fmt.Println(res)
}

func parseDo(input string) bool {
	return len(input) > 3 && input[0] == 'd' && input[1] == 'o' && input[2] == '(' && input[3] == ')'
}

func parseDont(input string) bool {
	return len(input) > 6 && input[0] == 'd' && input[1] == 'o' && input[2] == 'n' && input[3] == '\'' && input[4] == 't' && input[5] == '(' && input[6] == ')'
}

func solvePart2() {
	input, err := parseInput()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	shouldAddOperands := true
	var operands [][2]int
	for index := range input {
		if (shouldAddOperands && parseDont(input[index:])) || (!shouldAddOperands && parseDo(input[index:])) {
			shouldAddOperands = !shouldAddOperands
		}

		if shouldAddOperands {
			ops, err := getOperands(input[index:])
			if err != nil {
				continue
			}
			operands = append(operands, ops)
		}
	}

	var res int64 = 0
	for _, ops := range operands {
		res += int64(ops[0] * ops[1])
	}

	fmt.Println(res)
}

func main() {
	if os.Args[1] == "part1" {
		solvePart1()
	} else if os.Args[1] == "part2" {
		solvePart2()
	}
}
