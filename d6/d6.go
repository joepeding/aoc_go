package main

import (
	"fmt"
	"github.com/joepeding/aoc2021/common"
	"strconv"
	"strings"
)

func main() {
	var text = common.ReadLines("./d6/input")[0]
	fishes := parseInput(strings.Split(text, ","))
	fmt.Println(fishes)
	fmt.Println("Part one: ", partOne(fishes))
	fmt.Println("Part two: ", partTwo(fishes))
}

func parseInput(input []string) map[int]int64 {
	out := make(map[int]int64)
	for i := 0; i < 9; i++ {
		out[i] = int64(0)
	}
	for _, s := range input {
		d, err := strconv.Atoi(s)
		if err != nil { panic(err) }
		out[d]++
	}
	return out
}

func iterate(fishes map[int]int64) map[int]int64 {
	newFishes := fishes[0]
	for i := 1; i < 9; i++ {
		fishes[i-1] = fishes[i]
	}
	fishes[6] = fishes[6] + newFishes
	fishes[8] = newFishes
	return fishes
}

func partOne(fishes map[int]int64) int64 {
	for i := 0; i < 80; i++ {
		//fmt.Println(fishes)
		fishes = iterate(fishes)
	}
	out := int64(0)
	for i := 0; i < 9; i++ {
		out = out + fishes[i]
	}
	return out
}

func partTwo(fishes map[int]int64) int64 {
	for i := 0; i < 256-80; i++ {
		fmt.Println(fishes)
		fishes = iterate(fishes)
	}
	out := int64(0)
	for i := 0; i < 9; i++ {
		out = out + fishes[i]
	}
	return out
}
