package main

import (
	"fmt"
	"os"
	"strings"
)

const testData = `7-F7-
.FJ|7
SJLL7
|F--J
LJ.LJ`

const testData2 = `.F----7F7F7F7F-7....
.|F--7||||||||FJ....
.||.FJ||||||||L7....
FJL7L7LJLJ||LJ.L-7..
L--J.L7...LJS7F-7L7.
....F-J..F7FJ|L7L7L7
....L7.F7||L7|.L7L7|
.....|FJLJ|FJ|F7|.LJ
....FJL-7.||.||||...
....L---J.LJ.LJLJ...`

type Direction int

type Point struct {
	x int
	y int
}

type Vector struct {
	x     int
	y     int
	dir   Direction
	valid bool
}

type Maze struct {
	data  [][]byte
	start Point
}

const (
	UP = iota
	RIGHT
	DOWN
	LEFT
	Directions = 4
)

func surrounding(point Point, maze Maze) []Vector {
	surrounding := []Vector{}
	if point.y >= 0 && point.y < len(maze.data[0]) {
		if point.x > 0 {
			surrounding = append(surrounding, Vector{x: point.x - 1, y: point.y, dir: UP, valid: true})
		}
		if point.x < len(maze.data)-1 {
			surrounding = append(surrounding, Vector{x: point.x + 1, y: point.y, dir: DOWN, valid: true})
		}
	}
	if point.x >= 0 && point.x < len(maze.data) {
		if point.y > 0 {
			surrounding = append(surrounding, Vector{x: point.x, y: point.y - 1, dir: LEFT, valid: true})
		}
		if point.y < len(maze.data[0])-1 {
			surrounding = append(surrounding, Vector{x: point.x, y: point.y + 1, dir: RIGHT, valid: true})
		}
	}
	return surrounding
}

func pipeDirection(pipe byte) []Direction {
	switch pipe {
	case '|':
		return []Direction{UP, DOWN}
	case '-':
		return []Direction{LEFT, RIGHT}
	case 'F':
		return []Direction{DOWN, RIGHT}
	case '7':
		return []Direction{LEFT, DOWN}
	case 'L':
		return []Direction{UP, RIGHT}
	case 'J':
		return []Direction{LEFT, UP}
	}
	return []Direction{}
}

func direction(dir Direction) Point {
	switch dir {
	case UP:
		return Point{x: -1, y: 0}
	case RIGHT:
		return Point{x: 0, y: 1}
	case DOWN:
		return Point{x: 1, y: 0}
	case LEFT:
		return Point{x: 0, y: -1}
	}
	return Point{}
}

func opposite(dir Direction) Direction {
	switch dir {
	case UP:
		return DOWN
	case RIGHT:
		return LEFT
	case DOWN:
		return UP
	case LEFT:
		return RIGHT
	}
	return -1
}

func pprint(maze Maze) {
	for i := range maze.data {
		for j := range maze.data[i] {
			if maze.data[i][j] == '-' {
				fmt.Print("═")
			} else if maze.data[i][j] == '|' {
				fmt.Print("║")
			} else if maze.data[i][j] == 'F' {
				fmt.Print("╔")
			} else if maze.data[i][j] == '7' {
				fmt.Print("╗")
			} else if maze.data[i][j] == 'L' {
				fmt.Print("╚")
			} else if maze.data[i][j] == 'J' {
				fmt.Print("╝")
			} else {
				fmt.Print(string(maze.data[i][j]))
			}
		}
		fmt.Println()
	}
}

func followPipe(maze Maze, vector Vector) Vector {
	valid := false
	to := Vector{}
	directions := pipeDirection(maze.data[vector.x][vector.y])
	for _, pipe_dir := range directions {
		if pipe_dir == opposite(vector.dir) {
			valid = true
		} else {
			to = Vector{x: vector.x + direction(pipe_dir).x, y: vector.y + direction(pipe_dir).y, dir: pipe_dir, valid: true}
		}
	}
	if valid {
		return to
	} else {
		return Vector{valid: false}
	}
}

func parseData(data string) Maze {
	start := Point{x: 0, y: 0}
	rows := [][]byte{}
	for x, row := range strings.Split(data, "\n") {
		cols := []byte{}
		for y := range strings.Split(row, "") {
			char := row[y]
			cols = append(cols, char)
			if char == 'S' {
				start = Point{x: x, y: y}
			}
		}
		rows = append(rows, cols)
	}
	return Maze{data: rows, start: start}
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func amax(a [][]int) int {
	max := 0
	for i := range a {
		for j := range a[i] {
			if a[i][j] > max {
				max = a[i][j]
			}
		}
	}
	return max
}

func part1(maze Maze) int {
	dists := make([][]int, len(maze.data))

	for i := range dists {
		dists[i] = make([]int, len(maze.data[i]))
		for j := range dists[i] {
			dists[i][j] = -1
		}
	}
	for _, to := range surrounding(maze.start, maze) {
		step := 1
		dists[to.x][to.y] = step
		for {
			to = followPipe(maze, to)
			step++

			if maze.data[to.x][to.y] == 'S' {
				break
			}

			if !to.valid {
				break
			}

			if dists[to.x][to.y] == -1 {
				dists[to.x][to.y] = step
			} else {
				if step < dists[to.x][to.y] {
					dists[to.x][to.y] = step
				}
			}

			if step > 100000000000000 {
				fmt.Println("Too many steps")
				break
			}
		}

	}
	return amax(dists)
}

func part2(maze Maze) int {
	newmap := make([][]byte, len(maze.data))
	for i := range newmap {
		newmap[i] = make([]byte, len(maze.data[i]))
		for j := range newmap[i] {
			newmap[i][j] = ' '
		}
	}

	var final Vector
	var inside Direction
	for _, to := range surrounding(maze.start, maze) {
		complete := false
		last_dir := to.dir
		lefts := 0
		rights := 0
		final = to
		for {
			to = followPipe(maze, to)
			if !to.valid {
				break
			}

			if last_dir > to.dir || last_dir == LEFT && to.dir == UP {
				rights++
			} else if last_dir < to.dir || last_dir == UP && to.dir == LEFT {
				lefts++
			}
			newmap[to.x][to.y] = maze.data[to.x][to.y]
			if maze.data[to.x][to.y] == 'S' {
				complete = true
				if lefts > rights {
					inside = LEFT
				} else {
					inside = RIGHT
				}
				break
			}
		}
		if complete {
			newmap[final.x][final.y] = maze.data[final.x][final.y]
			break
		}
	}
	// pprint(Maze{data: newmap, start: maze.start})

	count := 0
	for to := final; newmap[to.x][to.y] != 'S'; to = followPipe(maze, to) {
		var inside_dir Direction
		if inside == LEFT {
			inside_dir = (to.dir + 1) % Directions
		} else {
			inside_dir = (to.dir - 1) % Directions
		}

		inside_point := Point{x: to.x + direction(inside_dir).x, y: to.y + direction(inside_dir).y}
		to_check := []Point{inside_point}
		for len(to_check) > 0 {
			pnt := to_check[0]
			to_check = to_check[1:]
			if pnt.x >= 0 && pnt.x < len(maze.data) && pnt.y >= 0 && pnt.y < len(maze.data[0]) {
				if newmap[pnt.x][pnt.y] == ' ' {
					count++
					newmap[pnt.x][pnt.y] = 'O'
					new_points := surrounding(pnt, maze)
					for _, new_point := range new_points {
						to_check = append(to_check, Point{x: new_point.x, y: new_point.y})
					}
				}
			}
		}
	}

	// pprint(Maze{data: newmap, start: maze.start})

	return count + 2
}

func test() {
	maze := parseData(testData2)
	// p1 := part1(maze)
	// fmt.Println("TEST Part 1: ", p1, "Expects: 8 Pass? ", p1 == 8)
	p2 := part2(maze)
	fmt.Println("TEST Part 2: ", p2, "Expects: 2 Pass? ", p2 == 2)
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

	maze := parseData(string(data))

	// p1 := part1(maze)
	// fmt.Println("Part 1: ", p1)
	p2 := part2(maze)
	fmt.Println("Part 2: ", p2)
}
