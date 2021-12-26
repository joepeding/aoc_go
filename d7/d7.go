package main

import (
	"fmt"
	"github.com/joepeding/aoc2021/common"
	"math"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var text = common.ReadLines("./d7/input")[0]
	positions := parseInput(text)
	//fmt.Println(positions)
	fmt.Println("Part one: ", partOne(positions))
	fmt.Println("Part two: ", partTwo(positions))
}

func parseInput(input string) []int {
	var out []int
	for _, s := range strings.Split(input, ",") {
		n, e := strconv.Atoi(s)
		if e != nil { panic (e) }
		out = append(out, n)
	}
	sort.Ints(out)
	return out
}

func fuelCost(pos int, positions []int) int {
	out := 0
	for _, p := range positions {
		out = out + int(math.Abs(float64(p-pos)))
	}
	return out
}

func partOne(positions []int) int {
	out := len(positions) * positions[len(positions) - 1]
	for i := 0; i < positions[len(positions) - 1]; i++ {
		cost := fuelCost(i, positions)
		//fmt.Println("Pos ", i, ": ", cost)
		if cost < out {
			out = cost
		} else {
			return out
		}
	}
	return out
}

func alternativeFuelCost(pos int, positions []int, fuelCosts []int) int {
	out := 0
	for _, p := range positions {
		dist := int(math.Abs(float64(p-pos)))
		out = out + fuelCosts[dist]
	}
	return out
}

func fuelCosts() []int {
	var out []int
	out = append(out, 0)
	for i := 1; i < 2000; i++ {
		out = append(out, out[i-1] + i)
	}
	return out
}

func partTwo(positions []int) int {
	var fuelCosts = fuelCosts()
	//fmt.Println(fuelCosts)
	out := len(positions) * positions[len(positions) - 1] * 10000
	for i := 0; i < positions[len(positions) - 1]; i++ {
		cost := alternativeFuelCost(i, positions, fuelCosts)
		fmt.Println("Pos ", i, ": ", cost)
		if cost < out {
			out = cost
		} else {
			//return out
		}
	}
	return out
}
