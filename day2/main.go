package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func parseInput() ([][]int, error) {
	var res [][]int

	file, err := os.Open("./input.txt")
	if err != nil {
		return res, err
	}

	bytes, err := io.ReadAll(file)
	if err != nil {
		return res, err
	}

	for _, line := range strings.Split(string(bytes), "\r\n") {
		var r []int
		for _, value := range strings.Split(line, " ") {
			i, err := strconv.Atoi(value)
			if err != nil {
				return res, err
			}
			r = append(r, i)
		}
		res = append(res, r)
	}

	return res, nil
}

func isSafeEntry(line []int) bool {
	var coef int
	for index := range line {
		if index == len(line)-1 {
			break
		}
		if index == 0 {
			if line[index] > line[index+1] {
				coef = 1
			} else if line[index] < line[index+1] {
				coef = -1
			}
		}
		if line[index] == line[index+1] {
			return false
		}
		if (line[index]-line[index+1])*coef < 0 {
			return false
		}
		if math.Abs(float64(line[index]-line[index+1])) > 3 {
			return false
		}
	}
	return true
}

func solvePart1() {
	input, err := parseInput()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	safe := 0
	for _, line := range input {
		if isSafeEntry(line) {
			safe += 1
		}
	}

	fmt.Println(safe)
}

func solvePart2() {
	input, err := parseInput()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	safe := 0
	for _, line := range input {
		isSafe := isSafeEntry(line)
		if isSafe {
			safe += 1
		} else {
			for index := range line {
				new := make([]int, len(line))
				copy(new, line)
				new = slices.Delete(new, index, index+1)
				if isSafeEntry(new) {
					safe += 1
					break
				}
			}
		}
	}

	fmt.Println(safe)
}

func main() {
	if os.Args[1] == "part1" {
		solvePart1()
	} else if os.Args[1] == "part2" {
		solvePart2()
	}
}
