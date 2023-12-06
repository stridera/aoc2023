package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func count_winners(line string) int {
	// fmt.Println("Line: ", line)
	card_data := strings.Split(line, ":")[1]
	numbers := strings.Split(card_data, "|")
	winning_numbers := strings.Split(numbers[0], " ")
	winners := 0
	for _, num := range strings.Split(numbers[1], " ") {
		if num == "" {
			continue
		}
		if slices.Contains(winning_numbers, num) {
			winners++
		}
	}
	return winners
}

func part1(data string) int {
	sum := 0
	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}
		winners := count_winners(line)
		if winners == 0 {
			continue
		}
		local_sum := 1
		for i := 1; i < winners; i++ {
			local_sum *= 2
		}

		sum += local_sum
	}
	return sum
}

func part2(data string) int {
	cards_processed := 0
	cards := strings.Split(string(data), "\n")
	if cards[len(cards)-1] == "" {
		cards = cards[:len(cards)-1]
	}
	cards_to_process := make([]int, len(cards))
	for i := range cards_to_process {
		cards_to_process[i] = i
	}

	for len(cards_to_process) > 0 {
		cards_processed++
		current_card := cards_to_process[0]
		cards_to_process = cards_to_process[1:]
		card := cards[current_card]
		winners := count_winners(card)
		if winners == 0 {
			continue
		}
		for i := 0; i < winners; i++ {
			cards_to_process = append(cards_to_process, current_card+1+i)
		}
	}
	return cards_processed
}

const testData = `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`

func test() {
	p1 := part1(testData)
	fmt.Println("TEST Part 1: ", p1, "Expects: 13 Pass? ", p1 == 13)
	p2 := part2(testData)
	fmt.Println("TEST Part 2: ", p2, "Expects: 30 Pass? ", p2 == 30)
}

func main() {
	test()
	// return

	// read file
	data, err := os.ReadFile("Day4/input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	sum := part1(string(data))
	fmt.Println("Part 1: ", sum)

	sum = part2(string(data))
	fmt.Println("Part 2: ", sum)
	// Wrong: 52997417
}
