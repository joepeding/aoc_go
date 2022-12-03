package main

import (
	"fmt"
	"github.com/joepeding/aoc_go/common"
	"sort"
)

func main() {
	var text = common.ReadLines("./2021/d10/input")
	fmt.Println("Part one: ", partOne(text))
	fmt.Println("Part two: ", partTwo(text))
}

func partOne(input []string) int {
	out := 0
	for _, s := range input {
		c := findIncorrectClosingChar(s)
		if c != 0 {
			out += charToScore(c)
		}
	}

	return out
}

func partTwo(input []string) int {
	filteredInput := make([]string, 0)
	for _, s := range input {
		if findIncorrectClosingChar(s) == 0 {
			filteredInput = append(filteredInput, s)
		}
	}

	closingScores := make([]int, 0)
	for _, s := range filteredInput {
		closingScores = append(closingScores, calculateClosingScore(s))
	}

	sort.Ints(closingScores)
	return closingScores[len(closingScores) / 2]
}

func findIncorrectClosingChar(s string) int32 {
	stack := make([]int32, 0)
	for i, c := range s {
		if i == 0 || isOpeningCharacter(c) {
			stack = append(stack, c)
			continue
		}
		oc := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if correctClosingChar(oc, c) {
			continue
		}
		return c
	}
	return 0
}

func calculateClosingScore(s string) int {
	stack := make([]int32, 0)
	for i, c := range s {
		if i == 0 || isOpeningCharacter(c) {
			stack = append(stack, c)
			continue
		}
		stack = stack[:len(stack)-1]
	}

	//Remaining chars on stack are all opening chars
	//Work backwards, because they're closed in reverse order
	score := 0
	for i := len(stack) - 1; i >= 0; i-- {
		score = score * 5 + openingCharToClosingScore(stack[i])
	}

	return score
}

func isOpeningCharacter(c int32) bool {
	for _, oc := range "([{<" {
		if c == oc {
			return true
		}
	}
	return false
}

func correctClosingChar(oc int32, cc int32) bool {
	if oc == '(' && cc == ')' {
		return true
	}
	if oc == '[' && cc == ']' {
		return true
	}
	if oc == '{' && cc == '}' {
		return true
	}
	if oc == '<' && cc == '>' {
		return true
	}
	return false
}

func charToScore(c int32) int {
	switch c {
		case ')': return 3
		case ']': return 57
		case '}': return 1197
		case '>': return 25137
	}
	return 0
}

func openingCharToClosingScore(c int32) int {
	switch c {
	case '(': return 1
	case '[': return 2
	case '{': return 3
	case '<': return 4
	}
	return 0
}

