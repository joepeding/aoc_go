package main

import (
	"fmt"
	"github.com/joepeding/aoc_go/common"
	"sort"
	"strconv"
	"strings"
)

type observation struct {
	patterns [][]string
	output [][]string
}

type mapping struct {
	zero []string
	one []string
	two []string
	three []string
	four []string
	five []string
	six []string
	seven []string
	eight []string
	nine []string
}

func main() {
	var text = common.ReadLines("./d8/input")
	observations := parseInput(text)
	fmt.Println("Part one: ", partOne(observations))
	fmt.Println("Part two: ", partTwo(observations))
}

func parseInput(input []string) []observation {
	var out []observation
	for _, s := range input {
		parts := strings.Split(s, " | ")

		var patterns [][]string
		for _, s := range strings.Split(parts[0], " ") {
			p := strings.Split(s, "")
			sort.Strings(p)
			patterns = append(patterns, p)
		}
		var output [][]string
		for _, s := range strings.Split(parts[1], " ") {
			o := strings.Split(s, "")
			sort.Strings(o)
			output = append(output, o)
		}

		out = append(out, observation{
			patterns: patterns,
			output: output,
		})
	}
	return out
}

func partOne(observations []observation) int {
	out := 0
	for _, o := range observations {
		for _, d := range o.output {
			switch len(d) {
				case 2,3,4,7: out++
			}
		}
	}
	return out
}

func complete(m mapping) bool {
	if len(m.zero) < 1 { return false }
	if len(m.one) < 1 { return false }
	if len(m.two) < 1 { return false }
	if len(m.three) < 1 { return false }
	if len(m.four) < 1 { return false }
	if len(m.five) < 1 { return false }
	if len(m.six) < 1 { return false }
	if len(m.seven) < 1 { return false }
	if len(m.eight) < 1 { return false }
	if len(m.nine) < 1 { return false }
	return true
}

func deduceMapping(patterns [][]string) mapping {
	m := mapping{}
	for {
		if complete(m) { break }
		for _, p := range patterns {
			// Skip known ones
			switch strings.Join(p, "") {
				case strings.Join(m.zero, ""),
					strings.Join(m.one, ""),
					strings.Join(m.two, ""),
					strings.Join(m.three, ""),
					strings.Join(m.four, ""),
					strings.Join(m.five, ""),
					strings.Join(m.six, ""),
					strings.Join(m.seven, ""),
					strings.Join(m.eight, ""),
					strings.Join(m.nine, ""): continue
			}
			// Fill easy ones
			switch len(p) {
				case 2: m.one = p
				case 3: m.seven = p
				case 4: m.four = p
				case 7: m.eight = p
			}
			// 3 is the only length-5 one which contains all of 1
			if len(p) == 5 && len(m.one) > 0 && containsAll(p, m.one) {
				m.three = p
				continue
			}
			// 6 is the only length-5 one which does not contain all of 1
			if len(p) == 6 && len(m.one) > 0 && !containsAll(p, m.one) {
				m.six = p
				continue
			}
			// 9 is th only length-6 one which contains all of 4
			if len(p) == 6 && len(m.four) > 0 && containsAll(p, m.four) {
				m.nine = p
				continue
			}
			// 5 is the only length 5 one which is contained by 6
			if len(p) == 5 && len(m.six) > 0 && containsAll(m.six, p) {
				m.five = p
				continue
			}
			// If three and five are filled, the remaining length-5 one must be 2
			if len(p) == 5 && len(m.five) > 0 && len(m.three) > 0 {
				m.two = p
				continue
			}
			// If six and nine are filled, the remaining length-6 one must b 0
			if len(p) == 6 && len(m.six) > 0 && len(m.nine) > 0 {
				m.zero = p
				continue
			}
		}
	}
	return m
}

func containsAll(set []string, subset []string) bool {
	contains := true
	OuterLoop:
	for _, c := range subset {
		for _, c2 := range set {
			if c == c2 {
				continue OuterLoop
			}
		}
		contains = false
	}
	return contains
}

func calcReading(m mapping, output [][]string) int {
	var ds []string
	for _, d := range output {
		switch strings.Join(d, "") {
			case strings.Join(m.zero, ""): ds = append(ds, "0")
			case strings.Join(m.one, ""): ds = append(ds, "1")
			case strings.Join(m.two, ""): ds = append(ds, "2")
			case strings.Join(m.three, ""): ds = append(ds, "3")
			case strings.Join(m.four, ""): ds = append(ds, "4")
			case strings.Join(m.five, ""): ds = append(ds, "5")
			case strings.Join(m.six, ""): ds = append(ds, "6")
			case strings.Join(m.seven, ""): ds = append(ds, "7")
			case strings.Join(m.eight, ""): ds = append(ds, "8")
			case strings.Join(m.nine, ""): ds = append(ds, "9")
		}
	}
	result, err := strconv.Atoi(strings.Join(ds, ""))
	if err != nil { panic(err) }
	fmt.Println(strings.Join(ds, ""))
	return result
}

func partTwo(observations []observation) int {
	out := 0
	for i,o := range observations {
		m := deduceMapping(o.patterns)
		r := calcReading(m, o.output)
		fmt.Println(i, ": ", m, " - ", o.output, " = ", r)
		out = out + r
	}
	return out
}

