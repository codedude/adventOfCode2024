/*
https://adventofcode.com/2024/day/13
Check day_13.jpeg for the maths :)
*/
package solutions

import (
	"aoc/internal"
	"strconv"
	"strings"
)

type Machine struct {
	A, B, P [2]int // [x,y]
}

type Machines = []*Machine

const (
	ACost int = 3
	BCost int = 1
)

func getInput13(part2 bool) Machines {
	data, _ := internal.ReadFileLinesWithBlanks("resources/day_13")
	machines := make(Machines, 0, 32)
	for i := 0; i < len(data); i += 4 {
		tmp := strings.Split(strings.Split(data[i], ": ")[1], ", ")
		ax, _ := strconv.Atoi(tmp[0][2:len(tmp[0])])
		ay, _ := strconv.Atoi(tmp[1][2:len(tmp[1])])
		tmp = strings.Split(strings.Split(data[i+1], ": ")[1], ", ")
		bx, _ := strconv.Atoi(tmp[0][2:len(tmp[0])])
		by, _ := strconv.Atoi(tmp[1][2:len(tmp[1])])
		tmp = strings.Split(strings.Split(data[i+2], ": ")[1], ", ")
		px, _ := strconv.Atoi(tmp[0][2:len(tmp[0])])
		py, _ := strconv.Atoi(tmp[1][2:len(tmp[1])])
		if part2 {
			px += 10000000000000
			py += 10000000000000
		}
		machines = append(machines, &Machine{A: [2]int{ax, ay}, B: [2]int{bx, by}, P: [2]int{px, py}})
	}
	return machines
}

func minInt(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func getGcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func priceReached(machine *Machine, aPressed, bPressed int) bool {
	return aPressed*machine.A[0]+bPressed*machine.B[0] == machine.P[0] && aPressed*machine.A[1]+bPressed*machine.B[1] == machine.P[1]
}

func solve13(machines Machines, part2 bool) int {
	totalMachine := []int{}
	for _, machine := range machines {
		// naive solution, iterate
		// max B button before getting over score
		// bPressed := minInt(machine.P[0]/machine.B[0], machine.P[1]/machine.B[1])
		// if !part2 { // naive solution dont work for large input
		// 	if bPressed == -1 { // no solution
		// 		break
		// 	}
		// 	// try to fill the gap with A (either X or Y it's the same)
		// 	aPressed := (machine.P[0] - (bPressed * machine.B[0])) / machine.A[0]
		// 	if priceReached(machine, aPressed, bPressed) { // solution
		// 		totalMachine = append(totalMachine, aPressed*ACost+bPressed*BCost)
		// 		break
		// 	}
		// 	bPressed -= 1 // next check
		// }
		bPressed := (machine.P[0]*machine.A[1] - machine.A[0]*machine.P[1]) / (-machine.A[0]*machine.B[1] + machine.A[1]*machine.B[0])
		aPressed := (machine.P[0] - machine.B[0]*bPressed) / machine.A[0]
		if priceReached(machine, aPressed, bPressed) { // solution
			totalMachine = append(totalMachine, aPressed*ACost+bPressed*BCost)
		}
	}
	total := 0
	for _, t := range totalMachine {
		total += t
	}
	return total
}

func Day13_1() int {
	data := getInput13(false)
	soluce := solve13(data, false)
	return soluce
}

func Day13_2() int {
	data := getInput13(true)
	soluce := solve13(data, true)
	return soluce
}
