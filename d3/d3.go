package main

import (
	"fmt"
	"github.com/joepeding/aoc2021/common"
	"math"
)

func main() {
	var text = common.ReadLines("./d3/input")
	fmt.Println("Part one: ", partOne(text))
	fmt.Println("Part two: ", partTwo 	(text))
}

func partOne(input []string) int {
	cf := make(map[int]int)
	for _, reading := range input {
		//fmt.Println("Analyzing: ", reading)
		for pos, char := range reading {
			//fmt.Println("-Analyzing pos: ", pos, " -> ", char)
			if char == 49 {
				//fmt.Println("--Incrementing")
				cf[pos]++
			}
		}
	}

	var g float64
	var e float64
	for pow := 0; pow < len(cf); pow++ {
		if cf[len(cf) - pow - 1] > len(input) / 2 {
			g = g + math.Pow(2, float64(pow))
		} else {
			e = e + math.Pow(2, float64(pow))
		}
	}
	fmt.Println("Gamma: ", g, " - Epsilon: ", e)
	return int(g * e)
}

func partTwo(input []string) int {
	fmt.Println("Determining O:")
	o := binaryStringToInt(filterCandidates(input, true))
	fmt.Println("Determining C:")
	c := binaryStringToInt(filterCandidates(input, false))

	fmt.Println("Oxygen: ", o, " - CO2: ", c)
	return o * c
}

func filterCandidates(input []string, mostCommon bool) string {
	fmt.Println("-Starting with: ", input)
	var filtered = input
	for pos := 0; pos < len(input[0]) && len(filtered) > 1; pos++ {
		filtered = filterCandidatesByCharacterInPos(filtered, pos, mostCommon)
		fmt.Println("-Remaining after Pos ", pos, ": ", filtered)
	}
	return filtered[0]
}

func filterCandidatesByCharacterInPos(input []string, pos int, mostCommon bool) []string {
	oc := float64(oneCountInPos(input, pos))

	var charToKeep int
	if mostCommon == true && oc >= float64(len(input)) / 2 { charToKeep = 49 }
	if mostCommon == true && oc < float64(len(input)) / 2 { charToKeep = 48 }
	if mostCommon == false && oc < float64(len(input)) / 2 { charToKeep = 49 }
	if mostCommon == false && oc >= float64(len(input)) / 2 { charToKeep = 48 }

	var filtered []string
	for _, s := range input {
		if int(s[pos]) == charToKeep {
			filtered = append(filtered, s)
		}
	}

	return filtered
}

func oneCountInPos(input []string, pos int) int {
	result := 0
	for _, reading := range input {
		if reading[pos] == 49 {
			result++
		}
	}
	return result
}

func binaryStringToInt(input string) int {
	var out float64
	for pow := 0; pow < len(input); pow++ {
		if input[len(input) - pow - 1] == 49 {
			out = out + math.Pow(2, float64(pow))
		}
	}
	return int(out)
}
