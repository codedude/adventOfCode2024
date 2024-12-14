/*
https://adventofcode.com/2024/day/14
*/
package solutions

import (
	"aoc/internal"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// height,width = example 11,7 / real input 103,101
const (
	roomWidth  int = 101
	roomHeight int = 103
	seconds    int = 100
)

type Robot struct {
	Id int    // Unique ID if needed
	P  [2]int // Position x,y
	V  [2]int // Velocity x,y (+x = right, +y = down)
}

type Robots = []*Robot

func getInput14(part2 bool) Robots {
	data, _ := internal.ReadFileLines("resources/day_14")
	robots := make(Robots, 0, 32)
	for i := 0; i < len(data); i++ {
		tmp := strings.Split(data[i], " ")
		tmpP := strings.Split(tmp[0][2:len(tmp[0])], ",")
		px, _ := strconv.Atoi(tmpP[0])
		py, _ := strconv.Atoi(tmpP[1])
		tmpV := strings.Split(tmp[1][2:len(tmp[1])], ",")
		vx, _ := strconv.Atoi(tmpV[0])
		vy, _ := strconv.Atoi(tmpV[1])
		robots = append(robots, &Robot{P: [2]int{px, py}, V: [2]int{vx, vy}, Id: i})
	}
	return robots
}

func (robot *Robot) Move(ticks int) {
	newPos := [2]int{(robot.P[0] + ticks*robot.V[0]) % roomWidth, (robot.P[1] + ticks*robot.V[1]) % roomHeight}
	if newPos[0] < 0 {
		newPos[0] += roomWidth
	}
	if newPos[1] < 0 {
		newPos[1] += roomHeight
	}
	robot.P[0] = newPos[0]
	robot.P[1] = newPos[1]
}

func countRobotQuadrant(robots Robots) (int, int, int, int) {
	midX, midY := roomWidth/2, roomHeight/2
	tl, tr, br, bl := 0, 0, 0, 0 // 4 quadrants
	for _, robot := range robots {
		if robot.P[0] == midX || robot.P[1] == midY { // skip mid
			continue
		}
		if robot.P[0] < midX && robot.P[1] < midY { //tl
			tl += 1
		} else if robot.P[0] > midX && robot.P[1] < midY { //tr
			tr += 1
		} else if robot.P[0] > midX && robot.P[1] > midY { //br
			br += 1
		} else if robot.P[0] < midX && robot.P[1] > midY { //bl
			bl += 1
		}
	}
	return tl, tr, br, bl
}

func solve14(robots Robots) int {
	for _, robot := range robots {
		robot.Move(seconds)
	}
	tl, tr, br, bl := countRobotQuadrant(robots)
	total := tl * tr * br * bl
	return total
}

// print in file to see the tree
func writeTree(robots Robots, second int, file *os.File) {
	room := make([][]int, roomHeight)
	for i := range room {
		room[i] = make([]int, roomWidth)
	}
	for _, robot := range robots {
		room[robot.P[1]][robot.P[0]] = 1
	}
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("%d - ", second))
	for i := range room {
		for j := range room[i] {
			if room[i][j] == 1 {
				builder.WriteByte('#')
			} else {
				builder.WriteByte('.')
			}
		}
		builder.WriteByte('\n')
	}
	builder.WriteByte('\n')
	file.WriteString(builder.String())
}

func solve14_2(robots Robots) int {
	var second int
	solution := -1

	file, err := os.OpenFile(
		"./day14_output",
		os.O_CREATE|os.O_TRUNC|os.O_WRONLY,
		0o764)
	if err != nil {
		return -1
	}
	defer file.Close()

	for second = 0; second < 10000; second++ { // found 10000 after 3 random input to have a reasonable upper bound
		for _, robot := range robots {
			robot.Move(1)
		}
		writeTree(robots, second, file)
	}
	return solution
}

func Day14_1() int {
	data := getInput14(false)
	soluce := solve14(data)
	return soluce
}

func Day14_2() int {
	data := getInput14(true)
	soluce := solve14_2(data)
	return soluce
}

func printRobots(robots Robots) {
	room := make([][]int, roomHeight)
	for i := range room {
		room[i] = make([]int, roomWidth)
	}
	for _, robot := range robots {
		room[robot.P[1]][robot.P[0]] += 1
	}
	for i := range room {
		for j := range room[i] {
			if room[i][j] == 0 {
				fmt.Printf(". ")
			} else {
				fmt.Printf("%d ", room[i][j])
			}
		}
		fmt.Println()
	}
}
