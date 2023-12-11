package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const testData = `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`

func parseData(data string) [][]int {
	num_data := [][]int{}
	for _, line := range strings.Split(data, "\n") {
		numbers := []int{}
		for _, num := range strings.Split(line, " ") {
			number, _ := strconv.Atoi(num)
			numbers = append(numbers, number)
		}
		num_data = append(num_data, numbers)
	}
	return num_data
}

func all(numbers []int, number int) bool {
	for _, num := range numbers {
		if num != number {
			return false
		}
	}
	return true
}

func findNextNumber(numbers []int) int {
	fmt.Println("FindNexNumber: ", numbers)
	for all(numbers, 0) {
		return 0
	}

	new_nums := []int{}
	for i := 1; i < len(numbers); i++ {
		new_nums = append(new_nums, numbers[i]-numbers[i-1])
	}
	return findNextNumber(new_nums) + numbers[len(numbers)-1]
}

func part1(data [][]int) int {
	sum := 0
	for _, numbers := range data {
		sum += findNextNumber(numbers)
	}
	return sum
}

func findPrevNumber(numbers []int) int {
	for all(numbers, 0) {
		return 0
	}

	new_nums := []int{}
	for i := 1; i < len(numbers); i++ {
		new_nums = append(new_nums, numbers[i]-numbers[i-1])
	}
	val := numbers[0] - findPrevNumber(new_nums)
	// fmt.Println("FindPrevNumber: ", val, numbers)
	return val
}

func part2(data [][]int) int {
	sum := 0
	for _, numbers := range data {
		sum += findPrevNumber(numbers)
	}
	return sum
}

func test() {
	data := parseData(testData)
	p1 := part1(data)
	fmt.Println("TEST Part 1: ", p1, "Expects: 114 Pass? ", p1 == 114)
	p2 := part2(data)
	fmt.Println("TEST Part 2: ", p2, "Expects: 2 Pass? ", p2 == 2)
}

func main() {
	test()
	// return

	// read file
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	numbers := parseData(string(data))
	p1 := part1(numbers)
	fmt.Println("Part 1: ", p1)
	p2 := part2(numbers)
	fmt.Println("Part 2: ", p2)
}
