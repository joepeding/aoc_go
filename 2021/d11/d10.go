package main

import (
	"fmt"
	"github.com/joepeding/aoc_go/common"
	"log"
	"strconv"
)

func main() {
	var text = common.ReadLines("./2021/d11/input")
	fmt.Println("Part one: ", partOne(text))
	fmt.Println("Part two: ", partTwo(text))
}

func partOne(input []string) int {
	out := 0

	grid := inputToGrid(input)
	for i := 0; i < 100; i++ {
		newGrid, newFlashes := round(grid)
		grid = newGrid
		out += newFlashes
	}
	fmt.Println(grid)

	return out
}

func partTwo(input []string) int {
	out := 0

	grid := inputToGrid(input)
	for {
		newGrid, newFlashes := round(grid)
		grid = newGrid
		out += 1
		if newFlashes > 99 {
			break
		}
	}

	return out
}

func inputToGrid(input []string) [][]int {
	grid := make([][]int, 0)
	for y := 0; y < 10; y++ {
		grid = append(grid, make([]int, 10))
		for x := 0; x < 10; x++ {
			num, err := strconv.Atoi(string(input[y][x]))
			if err != nil {
				log.Fatal(err)
			}
			grid[y][x] = num
		}
	}
	return grid
}

func round(input [][]int) ([][]int, int) {
	increased := increase(input)
	flashed := processFlash(increased)
	reset, flashes := reset(flashed)
	return reset, flashes
}

func increase(input [][]int) [][]int {
	output := input
	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			output[y][x] = output[y][x] + 1
		}
	}
	return output
}

func processFlash(input [][]int) [][]int {
	output := input

	// Create flash registry
	flashed := make([][]bool, 0)
	for y := 0; y < 10; y++ {
		flashed = append(flashed, make([]bool, 10))
	}

	for {
		stable := true

		for y := 0; y < 10; y++ {
			for x := 0; x < 10; x++ {
				if output[y][x] > 9 && flashed[y][x] == false {
					stable = false
					flashed[y][x] = true
					if y > 0 && x > 0 { output[y-1][x-1] = output[y-1][x-1] + 1}
					if y > 0          { output[y-1][x  ] = output[y-1][x  ] + 1}
					if y > 0 && x < 9 { output[y-1][x+1] = output[y-1][x+1] + 1}
					if          x > 0 { output[y  ][x-1] = output[y  ][x-1] + 1}
					if          x < 9 { output[y  ][x+1] = output[y  ][x+1] + 1}
					if y < 9 && x > 0 { output[y+1][x-1] = output[y+1][x-1] + 1}
					if y < 9          { output[y+1][x  ] = output[y+1][x  ] + 1}
					if y < 9 && x < 9 { output[y+1][x+1] = output[y+1][x+1] + 1}
				}
			}
		}

		if stable {
			break
		}
	}

	return output
}

func reset(input [][]int) ([][]int, int) {
	output := input
	flashes := 0

	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			if output[y][x] > 9 {
				output[y][x] = 0
				flashes += 1
			}
		}
	}

	return output, flashes
}
