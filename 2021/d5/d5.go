package main

import (
	"fmt"
	"github.com/joepeding/aoc_go/common"
	"math"
	"strconv"
	"strings"
)

type coord struct {
	x int
	y int
}

type line struct {
	a coord
	b coord
}

func main() {
	var text = common.ReadLines("./d5/input")
	lines, bottomright := parseInput(text)
	fmt.Println(lines)
	fmt.Println(bottomright)
	fmt.Println("Part one: ", partOne(lines, bottomright))
	fmt.Println("Part two: ", partTwo(lines, bottomright))
}

func parseInput(input []string) ([]line, coord) {
	bottomRight := coord{0,0}
	var output []line
	for _, s := range input {
		cs := strings.Split(s, " -> ")
		x1, errX1 := strconv.Atoi(strings.Split(cs[0], ",")[0])
		if errX1 != nil {panic(errX1)}
		y1, errY1 := strconv.Atoi(strings.Split(cs[0], ",")[1])
		if errY1 != nil {panic(errY1)}
		x2, errX2 := strconv.Atoi(strings.Split(cs[1], ",")[0])
		if errX2 != nil {panic(errX2)}
		y2, errY2 := strconv.Atoi(strings.Split(cs[1], ",")[1])
		if errY2 != nil {panic(errY2)}
		c1 := coord{x1, y1}
		c2 := coord{x2, y2}
		output = append(output, line{a: c1, b: c2})
		bottomRight.x = int(math.Max(float64(bottomRight.x), math.Max(float64(x1), float64(x2))))
		bottomRight.y = int(math.Max(float64(bottomRight.y), math.Max(float64(y1), float64(y2))))
	}
	return output, bottomRight
}

func noDiagonals(lines []line) []line {
	var output []line
	for _, l := range lines {
		if l.a.x == l.b.x || l.a.y == l.b.y {
			output = append(output, l)
		}
	}
	return output
}

func countCrossingLines(c coord, lines []line) int {
	out := 0
	for _, l := range lines {
		if l.a.x == l.b.x && c.x == l.a.x {
			// Vertical line
			if l.a.y > l.b.y {
				if c.y <= l.a.y && c.y >= l.b.y {
					out++
				}
			} else if l.b.y > l.a.y {
				if c.y <= l.b.y && c.y >= l.a.y {
					out++
				}
			} else {
				panic("0 length line")
			}
		} else if l.a.y == l.b.y && c.y == l.a.y {
			// Horizontal line
			if l.a.x > l.b.x {
				if c.x <= l.a.x && c.x >= l.b.x {
					out++
				}
			} else if l.b.x > l.a.x {
				if c.x <= l.b.x && c.x >= l.a.x {
					out++
				}
			} else {
				panic("0 length line")
			}
		} else {
			// Diagonal line
			var betweenX bool
			var betweenY bool
			var equidistant bool
			if l.a.x > l.b.x {
				betweenX = c.x >= l.b.x && c.x <= l.a.x
			} else {
				betweenX = c.x <= l.b.x && c.x >= l.a.x
			}
			if l.a.y > l.b.y {
				betweenY = c.y >= l.b.y && c.y <= l.a.y
			} else {
				betweenY = c.y <= l.b.y && c.y >= l.a.y
			}
			if math.Abs(float64(l.a.x - c.x)) == math.Abs(float64(l.a.y - c.y)) {
				equidistant = true
			}
			if betweenX && betweenY && equidistant {
				out++
			}
		}
	}
	return out
}

func partOne(lines []line, bottomRight coord) int {
	lines = noDiagonals(lines)
	out := 0
	for x := 0; x <= bottomRight.x; x++ {
		for y := 0; y <= bottomRight.y; y++ {
			if countCrossingLines(coord{x, y}, lines) >= 2 {
				out++
			}
		}
	}
	return out
}

func partTwo(lines []line, bottomRight coord) int {
	out := 0
	for x := 0; x <= bottomRight.x; x++ {
		for y := 0; y <= bottomRight.y; y++ {
			if countCrossingLines(coord{x, y}, lines) >= 2 {
				out++
			}
		}
	}
	return out
}
