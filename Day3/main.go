package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func hasSymbol(data []string, row int, start int, end int) bool {
	for i := max(0, row-1); i <= min(len(data)-1, row+1); i++ {
		for j := max(0, start-1); j <= min(len(data[i])-1, end+1); j++ {
			// Return true if there is a symbol (non digit or period) in the surrounding area
			if data[i][j] != '.' && !isDigit(data[i][j]) {
				return true
			}
		}
	}
	return false
}

func part1(data string) int {
	sum := 0

	lines := strings.Split(string(data), "\n")
	for row, str := range lines {
		// fmt.Println("************")
		// if row != 0 {
		// 	fmt.Println(row-1, lines[row-1])
		// }
		// fmt.Println(row, str)
		// if row != len(lines)-1 {
		// 	fmt.Println(row+1, lines[row+1])
		// }
		num := 0
		start := -1
		end := false
		for col, char := range str {
			if isDigit(byte(char)) {
				if start == -1 {
					start = col
				}
				num = num*10 + int(char-'0')
			} else if start != -1 {
				end = true
			}

			if start != -1 {
				if end || col == len(str)-1 {
					if hasSymbol(lines, row, start, col-1) {
						sum += num
						// fmt.Println("Symbol at [", start, ":", col, "] = ", num, "sum = ", sum)
					} else {
						// fmt.Println("No symbol at [", start, ":", col, "] = ", num, "sum = ", sum)
					}
					start = -1
					end = false
					num = 0
				}
			}
		}
	}

	return sum
}

func getNumber(str string, col int) int {
	start := col
	end := col
	for start >= 0 && isDigit(str[start]) {
		start--
	}
	for end < len(str) && isDigit(str[end]) {
		end++
	}
	num, _ := strconv.Atoi(str[start+1 : end])
	return num
}

func part2(data string) int {
	sum := 0

	lines := strings.Split(string(data), "\n")
	for row, str := range lines {
		for col, char := range str {
			if char == '*' {
				count := 0
				newNumber := true
				localProduct := 1
				// Need to check to see how many numbers are around it
				for i := max(0, row-1); i <= min(len(lines)-1, row+1); i++ {
					for j := max(0, col-1); j <= min(len(lines[i])-1, col+1); j++ {
						if isDigit(byte(lines[i][j])) {
							if newNumber {
								num := getNumber(lines[i], j)
								localProduct *= num
								// fmt.Println("Found number: ", num, " at [", i, ":", j, "] = ", localProduct)
								newNumber = false
								count++
							}
						} else {
							newNumber = true
						}
					}
					newNumber = true
				}
				if count == 2 {
					sum += localProduct
				}
				// if count > 0 {
				// 	fmt.Println("\n************")
				// 	if row != 0 {
				// 		fmt.Println(row-1, lines[row-1])
				// 	}
				// 	fmt.Println(row, str)
				// 	if row != len(lines)-1 {
				// 		fmt.Println(row+1, lines[row+1])
				// 	}
				// 	fmt.Print(row, " ")
				// 	for i := 0; i < len(lines[row]); i++ {
				// 		if i == col {
				// 			fmt.Print("^", localProduct, count)
				// 		} else {
				// 			fmt.Print(" ")
				// 		}
				// 	}
				// }
			}
		}
	}
	return sum
}

const testData = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

const testData2 = `467..114..
...*......
..354.633.
......#...
617.......
.....+.58.
..592.....
......755.
...$......
.664.598..`

func test() {
	p1 := part1(testData)
	fmt.Println("TEST Part 1: ", p1, "Expects: 4361 Pass? ", p1 == 4361)
	p2 := part2(testData)
	fmt.Println("TEST Part 2: ", p2, "Expects: 467835 Pass? ", p2 == 467835)
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

	sum := part1(string(data))
	fmt.Println("Part 1: ", sum)

	sum = part2(string(data))
	fmt.Println("Part 2: ", sum)
	// Wrong: 52997417
}
