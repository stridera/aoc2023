package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const testData = `Time:      7  15   30
Distance:  9  40  200`

type Race struct {
	time   int
	record int
}

func part1(data string) []Race {
	times := []int{}
	records := []int{}
	races := []Race{}

	lines := strings.Split(data, "\n")
	time_strings := strings.Split(lines[0], ":")
	for _, s := range strings.Split(time_strings[1], " ") {
		if s == "" {
			continue
		}
		i, _ := strconv.Atoi(s)
		times = append(times, i)
	}
	record_strings := strings.Split(lines[1], ":")
	for _, s := range strings.Split(record_strings[1], " ") {
		if s == "" {
			continue
		}
		i, _ := strconv.Atoi(s)
		records = append(records, i)
	}

	if len(times) != len(records) {
		fmt.Println("Error!  Times and Records don't match.")
	}
	for i := 0; i < len(times); i++ {
		races = append(races, Race{times[i], records[i]})
	}
	return races
}

func part2(data string) []Race {
	race := Race{}

	lines := strings.Split(data, "\n")
	time_strings := strings.Split(lines[0], ":")
	race.time, _ = strconv.Atoi(strings.Replace(time_strings[1], " ", "", -1))

	record_strings := strings.Split(lines[1], ":")
	race.record, _ = strconv.Atoi(strings.Replace(record_strings[1], " ", "", -1))

	return []Race{race}
}

func race(races []Race) int {
	val := 0

	for _, r := range races {
		local_sum := 0
		for i := 1; i < r.time; i++ {
			dist := i * (r.time - i)
			if dist > r.record {
				local_sum++
			}
		}
		if val == 0 {
			val = local_sum
		} else {
			val *= local_sum
		}
	}
	return val
}

func test() {
	races := part1(testData)
	p1 := race(races)
	fmt.Println("TEST Part 1: ", p1, "Expects: 288 Pass? ", p1 == 288)
	races = part2(testData)
	p2 := race(races)
	fmt.Println("TEST Part 2: ", p2, "Expects: 71503 Pass? ", p2 == 71503)
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

	races := part1(string(data))
	val := race(races)
	fmt.Println("Part 1: ", val)

	races = part2(string(data))
	val = race(races)
	fmt.Println("Part 2: ", val)

}
