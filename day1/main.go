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

func parseInput() ([]int, []int, error) {
	var left, right []int

	file, err := os.Open("./input.txt")
	if err != nil {
		return left, right, err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return left, right, err
	}

	for _, entry := range strings.Split(string(bytes), "\r\n") {
		values := strings.Split(entry, "   ")

		l, err := strconv.Atoi(values[0])
		if err != nil {
			return left, right, err
		}
		left = append(left, l)

		r, err := strconv.Atoi(values[1])
		if err != nil {
			return left, right, err
		}
		right = append(right, r)
	}

	return left, right, nil
}

func solvePart1() {
	left, right, err := parseInput()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	slices.Sort(left)
	slices.Sort(right)

	res := 0
	for index := range left {
		res += int(math.Abs(float64(left[index] - right[index])))
	}

	fmt.Println(res)
}

func countOccurances(list []int) map[int]int {
	m := make(map[int]int)
	for _, value := range list {
		_, ok := m[value]
		if !ok {
			m[value] = 1
		} else {
			m[value] += 1
		}
	}
	return m
}

func solvePart2() {
	left, right, err := parseInput()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	rightOccurances := countOccurances(right)

	res := 0
	for _, value := range left {
		occ, ok := rightOccurances[value]
		if ok {
			res += value * occ
		}
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
