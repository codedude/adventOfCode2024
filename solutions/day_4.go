/*
Day 4: https://adventofcode.com/2024/day/4
*/
package solutions

import (
	"aoc/internal"
	"sort"
	"strings"
)

// cellScore define the score of current cell for a pattern
type cellScore struct {
	r, b, br, bl int //direction right, bottom, bottom-right, bottom-left
}

func Day4_1() int {
	data, _ := internal.ReadFileLines("resources/day_4")

	patterns := []string{"XMAS", "SAMX"}

	// Helpers to manage the pivots table
	// currLine = the line being processed, lastLine = the line above
	var (
		currLine int = 1
		lastLine     = 0
	)
	total := 0

	// dynamic programming, pivots contains the current line + last line processed
	// lines of pivots are swaped every loop turn (in fact only ids are swap)
	pivots := make([][][]cellScore, 2)
	for i := range pivots {
		pivots[i] = make([][]cellScore, len(patterns))
		for j := range pivots[i] {
			pivots[i][j] = make([]cellScore, len(data[0]))
		}
	}
	for i := 0; i < len(data); i++ {
		total += computeScore(data[i], pivots, currLine, lastLine, patterns)
		currLine, lastLine = lastLine, currLine
	}

	return total
}

// computeScore calculate the score of the line line in pivots[currLine], based on pivots[lastLine]
// on first call, pivots[lastLine] is set to {0,0...}
// return the number of times it encounter a complete pattern
func computeScore(line string, pivots [][][]cellScore, currLine, lastLine int, pattern []string) int {
	patternFound := 0
	for ic, c := range line {
		for ip, p := range pattern {
			// c = current char, ic = index in pivots, p = current pattern, ip =  pattern id
			// 0 = not started, 1+ started
			cs := &pivots[currLine][ip][ic]               // cache the current cell score
			patIndex := strings.IndexByte(p, byte(c)) + 1 // current cell index in pattern (used for score)
			lastPatIndex := patIndex - 1                  // Last cell must contains this value to match pattern

			// New pattern starts
			if patIndex == 1 {
				cs.r, cs.b, cs.bl, cs.br = 1, 1, 1, 1
				continue
			} else {
				cs.r, cs.b, cs.bl, cs.br = 0, 0, 0, 0
			}
			// 1 < patIndex < 5 from here

			// UP (vert. top to bottom)
			if pivots[lastLine][ip][ic].b == lastPatIndex {
				cs.b = patIndex
				if patIndex == len(p) {
					patternFound += 1
				}
			} else {
				cs.b = 0
			}
			if ic > 0 {
				// LEFT (hor. left to right)
				if pivots[currLine][ip][ic-1].r == lastPatIndex {
					cs.r = patIndex
					if patIndex == len(p) {
						patternFound += 1
					}
				} else {
					cs.r = 0
				}
				// TOP LEFT (diag. top-left to bottom-right)
				if pivots[lastLine][ip][ic-1].br == lastPatIndex {
					cs.br = patIndex
					if patIndex == len(p) {
						patternFound += 1
					}
				} else {
					cs.br = 0
				}
			}
			if ic < len(line)-1 {
				// TOP RIGHT (diag. top-right to bottom-left)
				if pivots[lastLine][ip][ic+1].bl == lastPatIndex {
					cs.bl = patIndex
					if patIndex == len(p) {
						patternFound += 1
					}
				} else {
					cs.bl = 0
				}
			}

		}
	}
	return patternFound
}

func Day4_2() int {
	data, _ := internal.ReadFileLines("resources/day_4")
	patterns := []string{"MMSS", "MSMS", "SSMM", "SMSM"}
	sort.Strings(patterns)

	total := 0
	for i := 1; i < len(data)-1; i++ {
		for j := 1; j < len(data[0])-1; j++ {
			if data[i][j] == 'A' {
				foundPattern := string([]byte{data[i-1][j-1], data[i-1][j+1], data[i+1][j-1], data[i+1][j+1]})
				for _, v := range patterns {
					if v == foundPattern {
						total += 1
						break
					}
				}
			}
		}
	}
	return total
}
