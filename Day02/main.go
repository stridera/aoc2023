package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// The Elf would first like to know which games would have been possible if the bag
// contained only 12 red cubes, 13 green cubes, and 14 blue cubes?
func part1(line string) bool {
	for _, line := range strings.Split(line, ";") {
		for _, pulls := range strings.Split(line, ",") {
			pull := strings.Split(strings.TrimSpace(pulls), " ")
			count, _ := strconv.Atoi(string(pull[0]))
			color := pull[1]

			if color == "red" && count > 12 {
				return false
			}
			if color == "green" && count > 13 {
				return false
			}
			if color == "blue" && count > 14 {
				return false
			}
		}
	}
	return true
}

// In each game you played, what is the fewest number of cubes of each color that could have been in the
// bag to make the game possible?
// The power of a set of cubes is equal to the numbers of red, green, and blue cubes multiplied together.
func part2(line string) int {
	var red, green, blue int

	red = 0
	green = 0
	blue = 0

	for _, line := range strings.Split(line, ";") {
		for _, pulls := range strings.Split(line, ",") {
			pull := strings.Split(strings.TrimSpace(pulls), " ")
			count, _ := strconv.Atoi(string(pull[0]))
			color := pull[1]

			if color == "red" {
				red = max(red, count)
			} else if color == "green" {
				green = max(green, count)
			} else { // color == "blue"
				blue = max(blue, count)
			}
		}
	}
	return red * green * blue
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	sum := 0
	for id, line := range strings.Split(string(data), "\n") {
		if len(line) == 0 {
			continue
		}

		if part1(strings.Split(line, ":")[1]) {
			sum += id + 1
		}
	}
	fmt.Println("Part 1: ", sum)

	sum = 0
	for _, line := range strings.Split(string(data), "\n") {
		if len(line) == 0 {
			continue
		}

		sum += part2(strings.Split(line, ":")[1])
	}
	fmt.Println("Part 2: ", sum)
}
