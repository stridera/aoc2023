package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Maint struct {
	springs []rune
	groups  []int
}

func parseData(data string, part2 bool) []Maint {
	maints := []Maint{}
	for _, line := range strings.Split(data, "\n") {
		if line == "" {
			continue
		}
		parts := strings.Split(line, " ")
		springs := []rune(parts[0])
		groups := []int{}
		for _, s := range strings.Split(parts[1], ",") {
			if s == "" {
				continue
			}
			i, _ := strconv.Atoi(s)
			groups = append(groups, i)
		}

		maint := Maint{}
		maint.springs = springs
		maint.groups = groups
		if part2 {
			for i := 0; i < 4; i++ {
				maint.springs = append(maint.springs, '?')
				maint.springs = append(maint.springs, springs...)
				maint.groups = append(maint.groups, groups...)
			}
		}

		maints = append(maints, maint)
	}
	return maints
}

var cache = make(map[string]int)

func check(maint Maint, loc int, bloc int, step int) int {
	key := fmt.Sprintf("%d-%d-%d", loc, bloc, step)
	if v, ok := cache[key]; ok {
		return v
	}

	// fmt.Println("Check: ", string(maint.springs), ":", maint.groups, "Loc: ", loc, "BLoc: ", bloc, "Step: ", step)
	if loc == len(maint.springs) {
		if bloc == len(maint.groups) && step == 0 {
			return 1
		} else if bloc == len(maint.groups)-1 && maint.groups[bloc] == step {
			return 1
		} else {
			return 0
		}
	}

	sum := 0
	blocks := []rune{'.', '#'}
	for _, block := range blocks {
		// fmt.Println("Check: ", string(maint.springs), "Loc: ", loc, "Block: ", string(block), "BLoc: ", bloc, "Step: ", step)
		if maint.springs[loc] == block || maint.springs[loc] == '?' {
			if block == '.' && step == 0 {
				sum += check(maint, loc+1, bloc, 0)
			} else if block == '.' && step > 0 && bloc < len(maint.groups) && maint.groups[bloc] == step {
				sum += check(maint, loc+1, bloc+1, 0)
			} else if block == '#' {
				sum += check(maint, loc+1, bloc, step+1)
			}
		}
	}
	cache[key] = sum
	return sum
}

func calc(maints []Maint) int {
	sum := 0
	for _, maint := range maints {
		cache = make(map[string]int)
		score := check(maint, 0, 0, 0)
		// fmt.Println("Spring: ", string(maint.springs), "Groups: ", maint.groups, "Score: ", score)
		sum += score
	}

	return sum
}

const testData = `???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1`

func test() {
	maints := parseData(testData, false)
	p1 := calc(maints)
	fmt.Println("TEST Part 1: ", p1, "Expects: 21 Pass? ", p1 == 21)
	maints = parseData(testData, true)
	p2 := calc(maints)
	fmt.Println("TEST Part 2: ", p2, "Expects: 525152 Pass? ", p2 == 525152)
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

	maints := parseData(string(data), false)
	p1 := calc(maints)
	fmt.Println("Part 1: ", p1)

	maints = parseData(string(data), true)
	p2 := calc(maints)
	fmt.Println("Part 2: ", p2)
}
