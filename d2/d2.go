package main

import (
	"fmt"
	"github.com/joepeding/aoc2021/common"
	"strconv"
	"strings"
)

func main() {
	var text = common.ReadLines("./d2/input")
	fmt.Println("Part one: ", partOne(text))
	fmt.Println("Part one: ", partTwo(text))
}

func partOne(input []string) int {
	commands := sumCommands(input)
	return (commands["down"] - commands["up"]) * commands["forward"]
}

func sumCommands(input []string) map[string]int {
	commands := make(map[string]int)
	for _, command := range input {
		command := strings.Split(command, " ")
		dist, err := strconv.Atoi(command[1])
		if err != nil { panic(err) }
		commands[command[0]] = commands[command[0]] + dist
	}
	return commands
}

func partTwo(input []string) int {
	var h, d, a int

	for _, command := range input {
		h, d, a = processCommand(h, d, a, command)
	}

	return h * d
}

func processCommand(h, d, a int, command string) (int, int, int) {
	dir := strings.Split(command, " ")[0]
	dist, err := strconv.Atoi(strings.Split(command, " ")[1])
	if err != nil { panic(err) }

	if dir == "up" {
		a = a - dist
	} else if dir == "down" {
		a = a + dist
	} else {
		h = h + dist
		d = d + (a * dist)
	}

	return h, d, a
}

