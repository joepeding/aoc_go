package main

import (
	"fmt"
	"github.com/joepeding/aoc_go/common"
	"strconv"
	"strings"
)

type coord struct {
	x int
	y int
}

func main() {
	var text = common.ReadLines("./2021/d9/input")
	observations := parseInput(text)
	fmt.Println(observations)
	fmt.Println("Part one: ", partOne(observations))
	fmt.Println("Part two: ", partTwo(observations))
}

func parseInput(input []string) map[coord]int {
	out := make(map[coord]int)
	for y, xs := range input {
		for x, v := range strings.Split(xs, "") {
			val, e := strconv.Atoi(v)
			if e != nil { panic(e) }
			out[coord{x: x,	y: y}] = val
		}
	}
	return out
}

func isLow(p coord, m map[coord]int) bool {
	lowerNeighbours := 0
	l, lExists := m[coord{x: p.x - 1, y: p.y}]
	r, rExists := m[coord{x: p.x + 1, y: p.y}]
	t, tExists := m[coord{x: p.x, y: p.y - 1}]
	b, bExists := m[coord{x: p.x, y: p.y + 1}]

	if lExists && l <= m[p] { lowerNeighbours++ }
	if rExists && r <= m[p] { lowerNeighbours++ }
	if tExists && t <= m[p] { lowerNeighbours++ }
	if bExists && b <= m[p] { lowerNeighbours++ }

	return lowerNeighbours == 0
}

func findLows(m map[coord]int) []coord {
	var out []coord
	for c, _ := range m {
		if isLow(c, m) {
			out = append(out, c)
		}
	}
	return out
}

func partOne(observations map[coord]int) int {
	out := 0
	lows := findLows(observations)
	fmt.Println(lows)
	for _, c := range lows {
		out = out + observations[c] + 1
	}
	return out
}

func partTwo(observations map[coord]int) int {
	var basins [][]coord
	coordinatescan:
	for c,v := range observations {
		if v == 9 { continue coordinatescan}
		for _,b := range basins {
			for _, bc := range b {
				if c == bc { continue coordinatescan }
			}
		}

		newBasin := mapBasin(c, observations)
		basins = append(basins, newBasin)
		fmt.Println(newBasin)
	}

	var long1 []coord
	var long2 []coord
	var long3 []coord
	for _, b := range basins {
		if len(b) >= len(long1) {
			long3 = long2
			long2 = long1
			long1 = b
		} else if len(b) >= len(long2) {
			long3 = long2
			long2 = b
		} else if len(b) >= len(long3) {
			long3 = b
		}
	}

	fmt.Println(len(basins))
	return len(long1) * len(long2) * len(long3)
}

func mapBasin(p coord, m map[coord]int) []coord {
	var out []coord
	temp := make(map[coord]bool)
	temp[p] = false
	done := false

	for !done {
		done = true
		fmt.Println("Starting loop with len: ", len(temp), " / ", len(out))
		for c, checked := range temp {
			fmt.Println("Checking ", c, " - Checked: ", checked)
			if checked { continue }

			lCoord := coord{x: c.x - 1, y: c.y}
			rCoord := coord{x: c.x + 1, y: c.y}
			tCoord := coord{x: c.x, y: c.y - 1}
			bCoord := coord{x: c.x, y: c.y + 1}

			l, lExists := m[lCoord]
			r, rExists := m[rCoord]
			t, tExists := m[tCoord]
			b, bExists := m[bCoord]
			_, lAdded := temp[lCoord]
			_, rAdded := temp[rCoord]
			_, tAdded := temp[tCoord]
			_, bAdded := temp[bCoord]

			if lExists && l != 9 && !lAdded { temp[lCoord] = false; done = false; fmt.Println("Added", lCoord) }
			if rExists && r != 9 && !rAdded { temp[rCoord] = false; done = false; fmt.Println("Added", rCoord) }
			if tExists && t != 9 && !tAdded { temp[tCoord] = false; done = false; fmt.Println("Added", tCoord) }
			if bExists && b != 9 && !bAdded { temp[bCoord] = false; done = false; fmt.Println("Added", bCoord) }

			temp[c] = true
			out = append(out, c)
		}
	}
	return out
}
