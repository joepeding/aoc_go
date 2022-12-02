package main

import (
	"fmt"
	"github.com/joepeding/aoc_go/common"
	"strconv"
	"strings"
)

func main() {
	var text = common.ReadLines("./d4/input")
	seq, boards := parseInput(text)
	fmt.Println(seq, boards)
	fmt.Println("Part one: ", partOne(seq, boards))
	fmt.Println("Part two: ", partTwo(seq, boards))
}

func parseInput(input []string) ([]int, [][][]int) {
	seq := stringSliceToIntSlice(strings.Split(input[0], ","))

	rest := input[2:]
	var boards [][][]int
	var board [][]int
	for i, r := range rest {
		if (i+1) % 6 == 0 {
			boards = append(boards, board)
			board = make([][]int, 0)
		} else {
			board = append(board, stringSliceToIntSlice(strings.Split(r, " ")))
		}
	}
	boards = append(boards, board) // append last board

	return seq, boards
}

func stringSliceToIntSlice(input []string) []int {
	var output []int
	for _, s := range input {
		if s == "" { continue }
		i, e := strconv.Atoi(s)
		if (e != nil) { panic(e) }
		output = append(output, i)
	}
	return output
}

func partOne(seq []int, boards [][][]int) int {
	var draws []int
	for i, n := range seq {
		draws = seq[:i+1]
		fmt.Println("Draw ", i, ": ", n, "  -   Drawn so far: ", draws)
		for bn, b := range boards {
			fmt.Println("- Checking board ", bn)
			if hasBoardWonWithDraws(b, draws) {
				return scoreForBoard(b, draws)
			}
		}
	}
	return 0
}

func hasBoardWonWithDraws(board [][]int, draws []int) bool {
	markedBoard := markBoardWithDraws(board, draws)
	for _, r := range markedBoard {
		fmt.Println(r)
	}


rowcheck:
	for _, r := range markedBoard {
		for _, n := range r {
			if n != 0 {
				continue rowcheck
			}
		}
		return true
	}

colcheck:
	for c := 0; c < 5; c++ {
		for _, r := range markedBoard {
			if r[c] != 0 {
				continue colcheck
			}
		}
		return true
	}

	return false
}

func markBoardWithDraws(board [][]int, draws []int) [][]int {
	markedBoard := board
	for r := 0; r < 5; r++ {
		for c := 0; c < 5; c++ {
			for _, n := range draws {
				if markedBoard[r][c] == n {
					markedBoard[r][c] = 0
				}
			}
		}
	}
	return markedBoard
}

func scoreForBoard(board [][]int, draws []int) int {
	var sum int
	for _, r := range markBoardWithDraws(board, draws) {
		for _, n := range r {
			sum = sum + n
		}
	}

	return sum * draws[len(draws) - 1]
}

func partTwo(seq []int, boards [][][]int) int {
	var draws []int
	var rounds = 0
	for len(boards) > 1 {
		var remainingBoards = make([][][]int, 0)
		draws = seq[:rounds+1]
		fmt.Println("Draw ", rounds, ": ", seq[rounds], "  -   Drawn so far: ", draws)
		for bn, b := range boards {
			fmt.Println("- Checking board ", bn)
			if hasBoardWonWithDraws(b, draws) == false {
				remainingBoards = append(remainingBoards, b)
			}
		}
		boards = remainingBoards
		rounds++
	}
	for hasBoardWonWithDraws(boards[0], draws) == false {
		rounds++
		draws = seq[:rounds + 1]
		fmt.Println("-")
	}
	return scoreForBoard(boards[0], draws)
}
