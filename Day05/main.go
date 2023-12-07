package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const testData = `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`

type Locs struct {
	dest_start   int
	source_start int
	size         int
}

type Maps struct {
	from string
	to   string
	locs []Locs
}

func parse_data(data string) ([]string, []Maps) {
	sections := strings.Split(data, "\n\n")
	seeds := strings.Split(strings.Split(sections[0], ": ")[1], " ")
	maps := []Maps{}
	for sect := 1; sect < len(sections); sect++ {
		lines := strings.Split(sections[sect], "\n")
		map_line := strings.Split(strings.Split(lines[0], " ")[0], "-")
		from := map_line[0]
		to := map_line[2]
		locs := []Locs{}

		for i := 1; i < len(lines); i++ {
			if lines[i] == "" {
				continue
			}
			numbers := strings.Split(lines[i], " ")
			dest_start, _ := strconv.Atoi(numbers[0])
			source_start, _ := strconv.Atoi(numbers[1])
			size, _ := strconv.Atoi(numbers[2])
			locs = append(locs, Locs{dest_start, source_start, size})
		}
		maps = append(maps, Maps{from, to, locs})
	}
	return seeds, maps
}

func follow_map(val int, from string, destination string, maps []Maps) int {
	// fmt.Print(from, " ", val, " ")
	if from == destination {
		return val
	}
	for _, m := range maps {
		if from == m.from {
			for _, l := range m.locs {
				if val >= l.source_start && val < l.source_start+l.size {
					val = l.dest_start + (val - l.source_start)
					return follow_map(val, m.to, destination, maps)
				}
			}
			return follow_map(val, m.to, destination, maps)
		}
	}
	fmt.Println("ERROR: No map found for ", from, destination)
	return 0
}

func part1(seeds []string, maps []Maps) int {
	lowest := int(^uint(0) >> 1)
	for _, seed := range seeds {
		seed, _ := strconv.Atoi(seed)
		location := follow_map(seed, "seed", "location", maps)
		if location < lowest {
			lowest = location
		}
	}
	return lowest
}

func part2(seeds []string, maps []Maps) int {
	lowest := int(^uint(0) >> 1)
	start := -1
	for _, seed := range seeds {
		seed, _ := strconv.Atoi(seed)
		if start == -1 {
			start = seed
		} else {
			for s := start; s < start+seed; s++ {
				val := follow_map(s, "seed", "location", maps)
				if val < lowest {
					lowest = val
				}
			}
			start = -1
		}
	}
	return lowest
}

func test() {
	seeds, maps := parse_data(testData)

	p1 := part1(seeds, maps)
	fmt.Println("TEST Part 1: ", p1, "Expects: 35 Pass? ", p1 == 35)
	p2 := part2(seeds, maps)
	fmt.Println("TEST Part 2: ", p2, "Expects: 46 Pass? ", p2 == 46)
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
	seeds, maps := parse_data(string(data))

	val := part1(seeds, maps)
	fmt.Println("Part 1: ", val)

	val = part2(seeds, maps)
	fmt.Println("Part 2: ", val)
}
