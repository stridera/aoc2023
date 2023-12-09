package main

import (
	"fmt"
	"os"
	"strings"
)

const testData = `LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)`

const testData2 = `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`

func parseData(data string) (string, map[string][]string) {
	split := strings.Split(data, "\n\n")
	path := split[0]
	maps := map[string][]string{}
	for _, line := range strings.Split(split[1], "\n") {
		if line == "" {
			continue
		}
		parts := strings.Split(line, " = ")
		maps[parts[0]] = strings.Split(parts[1][1:len(parts[1])-1], ", ")
	}
	return path, maps
}

func recursive_walk(location string, part1 bool, path string, maps map[string][]string) int {
	if at_destination(part1, location) {
		return 0
	}
	next_step := path[0]
	path = path[1:] + string(next_step)
	if next_step == 'L' {
		return recursive_walk(maps[location][0], part1, path, maps) + 1
	} else if next_step == 'R' {
		return recursive_walk(maps[location][1], part1, path, maps) + 1
	}

	return 0
}

func at_destination(part1 bool, location string) bool {
	if part1 {
		if location == "ZZZ" {
			return true
		}
		return false
	} else {
		if location[2] != 'Z' {
			return false
		}
		return true
	}
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(nums []int) int {
	result := nums[0]
	for i := 1; i < len(nums); i++ {
		result = result * nums[i] / gcd(result, nums[i])
	}
	return result
}

func part1(path string, maps map[string][]string) int {
	return recursive_walk("AAA", true, path, maps)
}

func part2(path string, maps map[string][]string) int {
	locations := []string{}
	for location := range maps {
		if location[2] == 'A' {
			locations = append(locations, location)
		}
	}

	lengths := []int{}
	for _, location := range locations {
		lengths = append(lengths, recursive_walk(location, false, path, maps))
	}
	return lcm(lengths)
}

func test() {
	path, maps := parseData(testData)
	p1 := part1(path, maps)
	fmt.Println("TEST Part 1: ", p1, "Expects: 6 Pass? ", p1 == 6)
	path, maps = parseData(testData2)
	p2 := part2(path, maps)
	fmt.Println("TEST Part 2: ", p2, "Expects: 6 Pass? ", p2 == 6)
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

	path, maps := parseData(string(data))
	p1 := part1(path, maps)
	fmt.Println("Part 1: ", p1)
	p2 := part2(path, maps)
	fmt.Println("Part 2: ", p2)
}
