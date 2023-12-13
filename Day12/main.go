package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const testData = `???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1`

type Maint struct {
	springs []rune
	groups  []int
}

func parseData(data string) []Maint {
	maints := []Maint{}
	for _, line := range strings.Split(data, "\n") {
		if line == "" {
			continue
		}
		parts := strings.Split(line, " ")
		maint := Maint{}
		maint.springs = []rune(parts[0])
		maint.groups = []int{}
		for _, s := range strings.Split(parts[1], ",") {
			if s == "" {
				continue
			}
			i, _ := strconv.Atoi(s)
			maint.groups = append(maint.groups, i)
		}
		maints = append(maints, maint)
	}
	return maints
}

func createPerms(springs []rune) [][]rune {
	fmt.Println("Create Perms: ", string(springs))

	perms := [][]rune{}
	for _, spring := range springs {
		new_perms := [][]rune{}
		if spring == '?' {
			if len(perms) == 0 {
				new_perms = append(new_perms, []rune{'.'})
				new_perms = append(new_perms, []rune{'#'})
			} else {
				for _, p := range perms {
					new_perms = append(new_perms, [][]rune{append(p, 'd'), append(p, 'h')}...)
				}
			}
		} else {
			if len(perms) == 0 {
				new_perms = append(new_perms, []rune{spring})
			} else {
				for _, p := range perms {
					new_perms = append(new_perms, append(p, 'E'))
				}
			}
		}
		fmt.Println("Spring: ", string(spring), "Perms: ", len(new_perms))
		for _, p := range new_perms {
			fmt.Println(string(p))
		}
		perms = [][]rune{}
		for _, p := range new_perms {
			perms = append(perms, p)
		}
	}
	// Remove Dups
	seen := map[string]bool{}
	for _, perm := range perms {
		seen[string(perm)] = true
	}
	perms = [][]rune{}
	for k := range seen {
		perms = append(perms, []rune(k))
	}
	// fmt.Println("Perms: ", len(perms))
	for _, p := range perms {
		fmt.Println(string(p))
	}
	return perms

}

func isValid(springs []rune, groups []int) bool {
	for {
		if len(springs) == 0 && len(groups) == 0 {
			return true
		} else if len(springs) == 0 {
			return false
		}

		if springs[0] == '.' {
			springs = springs[1:]
		} else if springs[0] == '#' {
			group := groups[0]
			if len(springs) < group {
				return false
			}
			for i := 0; i < group; i++ {
				if springs[i] != '#' {
					return false
				}
			}
			springs = springs[group:]
			groups = groups[1:]
			if len(springs) == 0 && len(groups) == 0 {
				return true
			} else if len(springs) > 1 && springs[0] == '#' {
				return false
			}

		}
	}
}

func part1(maints []Maint) int {
	createPerms([]rune("#?#?#?#"))
	return 0

	sum := 0
	for _, maint := range maints {
		fmt.Println("Spring: ", string(maint.springs), "Groups: ", maint.groups)
		perms := createPerms(maint.springs)
		for _, perm := range perms {
			// fmt.Println("- ", string(perm))

			if isValid(perm, maint.groups) {
				fmt.Println(string(perm), maint.groups, " is valid")
				sum++
			}
		}
	}

	return sum
}

func part2(maints []Maint) int {
	return 0
}

func test() {
	maints := parseData(testData)
	p1 := part1(maints)
	fmt.Println("TEST Part 1: ", p1, "Expects: 21 Pass? ", p1 == 21)
	p2 := part2(maints)
	fmt.Println("TEST Part 2: ", p2, "Expects: 1030 Pass? ", p2 == 1030)
}

func main() {
	test()
	return

	// read file
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	maints := parseData(string(data))
	p1 := part1(maints)
	fmt.Println("Part 1: ", p1)

	p2 := part2(maints)
	fmt.Println("Part 2: ", p2)
}
