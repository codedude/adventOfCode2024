/*
https://adventofcode.com/2024/day/15
*/
package solutions

import (
	"aoc/internal"
	"fmt"
)

type WarehouseCell int

const (
	WarehouseFree WarehouseCell = iota
	WarehouseBox
	WarehouseBoxL
	WarehouseBoxR
	WarehouseWall
)

type Warehouse struct {
	Room  [][]WarehouseCell // 2d array of the room
	Robot [2]int            // Position of the robot
	Moves []WalkDirection   // List of movements the robot will try to do
}

func getInput15(part2 bool) *Warehouse {
	data, _ := internal.ReadFileLinesWithBlanks("resources/day_15")
	warehouse := Warehouse{}
	dataPart := 0
	for i, line := range data {
		if len(line) == 0 {
			dataPart = i + 1 // when movements list starts
			continue
		}
		if dataPart == 0 { // room
			roomLine := []WarehouseCell{}
			for j, cell := range data[i] {
				var cellType WarehouseCell
				if cell == '#' {
					cellType = WarehouseWall
				} else if cell == 'O' {
					cellType = WarehouseBox
				} else if cell == '@' {
					cellType = WarehouseFree
					if part2 {
						warehouse.Robot = [2]int{i, j * 2}
					} else {
						warehouse.Robot = [2]int{i, j}
					}
				} else {
					cellType = WarehouseFree
				}
				if part2 {
					if cellType == WarehouseBox {
						roomLine = append(roomLine, WarehouseBoxL)
						roomLine = append(roomLine, WarehouseBoxR)
					} else {
						roomLine = append(roomLine, cellType)
						roomLine = append(roomLine, cellType)
					}
				} else {
					roomLine = append(roomLine, cellType)
				}
			}
			warehouse.Room = append(warehouse.Room, roomLine)
		} else { // movements
			for _, move := range data[i] {
				var moveDir WalkDirection
				if move == '^' {
					moveDir = UP
				} else if move == '>' {
					moveDir = RIGHT
				} else if move == 'v' {
					moveDir = DOWN
				} else {
					moveDir = LEFT
				}
				warehouse.Moves = append(warehouse.Moves, moveDir)
			}
		}
	}
	return &warehouse
}

func getNextPos(pos, move [2]int) [2]int {
	return [2]int{pos[0] + move[0], pos[1] + move[1]}
}

func recMoveBlock(warehouse *Warehouse, direction WalkDirection, pos [2]int, boxesToMove *[][2]int) bool {
	nextPos := getNextPos(pos, getNextMove(direction))
	cell := warehouse.Room[pos[0]][pos[1]]
	if cell == WarehouseBoxL { // [, j+1 for ]
		*boxesToMove = append(*boxesToMove, pos)
		return recMoveBlock(warehouse, direction, nextPos, boxesToMove) && recMoveBlock(warehouse, direction, [2]int{nextPos[0], nextPos[1] + 1}, boxesToMove)
	} else if cell == WarehouseBoxR { // ], j-1 for []
		*boxesToMove = append(*boxesToMove, [2]int{pos[0], pos[1] - 1})
		return recMoveBlock(warehouse, direction, nextPos, boxesToMove) && recMoveBlock(warehouse, direction, [2]int{nextPos[0], nextPos[1] - 1}, boxesToMove)
	} else if cell == WarehouseFree {
		return true
	} else {
		return false
	}
}

func moveRobot2(warehouse *Warehouse, direction WalkDirection, nextPos [2]int) {
	if direction == RIGHT || direction == LEFT {
		// find first free space, move all by one
		freePos := nextPos
		foundFree := false
		for {
			freePos = getNextPos(freePos, getNextMove(direction)) // pos after the box
			if warehouse.Room[freePos[0]][freePos[1]] == WarehouseWall {
				break // wall so cant move
			} else if warehouse.Room[freePos[0]][freePos[1]] == WarehouseFree {
				foundFree = true
				break
			} else {
				continue // box can be move together
			}
		}
		if foundFree {
			var way int
			if direction == RIGHT {
				way = -1
			} else {
				way = 1
			}
			// move all by one, from free pos found to robot
			for p := freePos[1]; p != nextPos[1]; p += way {
				warehouse.Room[freePos[0]][p], warehouse.Room[freePos[0]][p+way] = warehouse.Room[freePos[0]][p+way], warehouse.Room[freePos[0]][p]
			}
			warehouse.Robot = nextPos
		}
	} else { // UP/DOWN
		boxesToMove := make([][2]int, 0, 4)
		if recMoveBlock(warehouse, direction, nextPos, &boxesToMove) {
			var way int
			if direction == UP {
				way = -1
			} else {
				way = 1
			}
			for _, boxToMove := range boxesToMove {
				warehouse.Room[boxToMove[0]][boxToMove[1]] = WarehouseFree
				warehouse.Room[boxToMove[0]][boxToMove[1]+1] = WarehouseFree
			}
			for _, boxToMove := range boxesToMove {
				warehouse.Room[boxToMove[0]+way][boxToMove[1]] = WarehouseBoxL
				warehouse.Room[boxToMove[0]+way][boxToMove[1]+1] = WarehouseBoxR
			}
			warehouse.Robot = [2]int{warehouse.Robot[0] + way, warehouse.Robot[1]}
		}
	}
}

func moveRobot(warehouse *Warehouse, direction WalkDirection, part2 bool) {
	nextPos := getNextPos(warehouse.Robot, getNextMove(direction))
	if warehouse.Room[nextPos[0]][nextPos[1]] == WarehouseWall {
		return // wall = skip
	} else if warehouse.Room[nextPos[0]][nextPos[1]] == WarehouseFree {
		warehouse.Robot = nextPos
		return // move it
	} else { // try to move block
		if part2 {
			moveRobot2(warehouse, direction, nextPos)
		} else {
			secNextPos := nextPos
			foundFree := false
			for {
				secNextPos = getNextPos(secNextPos, getNextMove(direction)) // pos after the box
				if warehouse.Room[secNextPos[0]][secNextPos[1]] == WarehouseWall {
					break // wall so cant move
				} else if warehouse.Room[secNextPos[0]][secNextPos[1]] == WarehouseBox {
					continue // box can be move together
				} else { // free cell
					foundFree = true
					break
				}
			}
			if foundFree {
				warehouse.Room[secNextPos[0]][secNextPos[1]] = WarehouseBox
				warehouse.Room[nextPos[0]][nextPos[1]] = WarehouseFree
				warehouse.Robot = nextPos
			}
		}
	}
}

func countBoxGps(warehouse *Warehouse, part2 bool) int {
	total := 0
	for i, line := range warehouse.Room {
		for j, cell := range line {
			if part2 {
				if cell == WarehouseBoxL {
					total += 100*i + j
				}
			} else {
				if cell == WarehouseBox {
					total += 100*i + j
				}
			}
		}
	}
	return total
}

func printWarehouse(warehouse *Warehouse) {
	for _, line := range warehouse.Room {
		fmt.Println(line)
	}
	fmt.Println()
}

func solve15(warehouse *Warehouse, part2 bool) int {
	for _, move := range warehouse.Moves {
		moveRobot(warehouse, move, part2)
	}
	score := countBoxGps(warehouse, part2)
	return score
}

func Day15_1() int {
	data := getInput15(false)
	soluce := solve15(data, false)
	return soluce
}

func Day15_2() int {
	data := getInput15(true)
	soluce := solve15(data, true)
	return soluce
}
