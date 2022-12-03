package main

import (
	"fmt"
	"github.com/joepeding/aoc_go/common"
	"strconv"
)

func main() {
	var text = common.ReadLines("./2021/d1/input")
	fmt.Println("Part one: ", partOne(text))
	fmt.Println("Part one: ", partTwo(text))
}

func partOne(input []string) int {
	var increases = 0
	for index, line := range input {
		if (index == 0) { continue }
		var current, errP = strconv.ParseInt(line, 10, 16)
		var previous, errC = strconv.ParseInt(input[index - 1], 10, 16)
		if (errP != nil) {
			panic(errP)
		}
		if (errC != nil) {
			panic(errC)
		}
		if (current > previous) {
			increases++
		}
	}
	return increases
}

func partTwo(input []string) int {
	var increases = 0
	for index, line := range input {
		if (index < 3) { continue }
		var current, _ = strconv.ParseInt(line, 10, 16)
		var minus1, _ = strconv.ParseInt(input[index - 1], 10, 16)
		var minus2, _ = strconv.ParseInt(input[index - 2], 10, 16)
		var minus3, _ = strconv.ParseInt(input[index - 3], 10, 16)
		if current + minus1 + minus2 > minus1 + minus2 + minus3 {
			increases++
		}
	}
	return increases
}
