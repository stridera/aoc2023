package main

import (
	"fmt"
	"os"
	"strings"
)

const testData = `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`

type Point struct {
	x int
	y int
}

func pprint(points []Point) {
	max_x := 0
	max_y := 0
	for _, point := range points {
		if point.x > max_x {
			max_x = point.x
		}
		if point.y > max_y {
			max_y = point.y
		}
	}

	newmap := make([][]byte, max_y+1)
	for i := range newmap {
		newmap[i] = make([]byte, max_x+1)
		for j := range newmap[i] {
			newmap[i][j] = '.'
		}
	}
	for _, point := range points {
		newmap[point.y][point.x] = '#'
	}
	for _, line := range newmap {
		fmt.Println(string(line))
	}
}

func parseData(data string, expansion int) []Point {
	points := []Point{}
	lines := strings.Split(data, "\n")
	for y, line := range lines {
		for x, char := range line {
			if char == '#' {
				points = append(points, Point{x: x, y: y})
			}
		}
	}

	// pprint(points)

	star_map_x := make(map[int]int)
	star_map_y := make(map[int]int)
	max_x := 0
	max_y := 0
	for _, point := range points {
		star_map_x[point.x]++
		star_map_y[point.y]++
		if point.x > max_x {
			max_x = point.x
		}
		if point.y > max_y {
			max_y = point.y
		}
	}

	for x := max_x; x > 0; x-- {
		if star_map_x[x] == 0 {
			for p := range points {
				if points[p].x > x {
					points[p].x += expansion
				}
			}
		}
	}
	for y := max_y; y > 0; y-- {
		if star_map_y[y] == 0 {
			for p := range points {
				if points[p].y > y {
					points[p].y += expansion
				}
			}
		}
	}

	// pprint(points)

	return points
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func pathsum(points []Point) int {
	sum := 0
	for i := 0; i < len(points); i++ {
		p1 := points[i]
		for j := i; j < len(points); j++ {
			p2 := points[j]
			if p1 == p2 {
				continue
			}
			dist := abs(p1.x-p2.x) + abs(p1.y-p2.y)
			// fmt.Println(i+1, j+1, p1, p2, dist)
			sum += dist
		}
	}
	return sum
}

func test() {
	sky := parseData(testData, 1)
	p1 := pathsum(sky)
	fmt.Println("TEST Part 1: ", p1, "Expects: 374 Pass? ", p1 == 374)
	sky = parseData(testData, 10-1)
	p2 := pathsum(sky)
	fmt.Println("TEST Part 2: ", p2, "Expects: 1030 Pass? ", p2 == 1030)
	sky = parseData(testData, 100-1)
	p2 = pathsum(sky)
	fmt.Println("TEST Part 2: ", p2, "Expects: 8410 Pass? ", p2 == 8410)
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

	sky := parseData(string(data), 1)
	p1 := pathsum(sky)
	fmt.Println("Part 1: ", p1)

	sky = parseData(string(data), 1_000_000-1)
	p2 := pathsum(sky)
	fmt.Println("Part 2: ", p2)
}
