package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const numbers = "0123456789"

func part1(line string) int {
	first_index := strings.IndexAny(line, numbers)
	last_index := strings.LastIndexAny(line, numbers)
	if first_index == -1 || last_index == -1 {
		return 0
	}
	num, _ := strconv.Atoi(string(line[first_index]) + string(line[last_index]))
	return num
}

func part2(line string) int {
	words := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	first_index := strings.IndexAny(line, numbers)
	last_index := strings.LastIndexAny(line, numbers)
	if first_index == -1 || last_index == -1 {
		return 0
	}
	first_digit, _ := strconv.Atoi(string(line[first_index]))
	last_digit, _ := strconv.Atoi(string(line[last_index]))
	for i, word := range words {
		if strings.Contains(line, word) {
			if strings.Index(line, word) < first_index {
				first_index = strings.Index(line, word)
				first_digit = i
			}
			if strings.LastIndex(line, word) > last_index {
				last_index = strings.LastIndex(line, word)
				last_digit = i
			}
		}
	}
	return first_digit*10 + last_digit
}

func main() {
	// read file
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	sum := 0
	for _, line := range strings.Split(string(data), "\n") {
		sum += part1(line)
	}
	fmt.Println("Part 1: ", sum)

	sum = 0
	for _, line := range strings.Split(string(data), "\n") {
		sum += part2(line)
	}
	fmt.Println("Part 2: ", sum)
}
