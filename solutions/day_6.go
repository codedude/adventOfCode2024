/*
https://adventofcode.com/2024/day/6
*/
package solutions

import (
	"aoc/internal"
)

type MapCell int

const (
	EMPTY MapCell = iota
	WALL
	SEEN
)

type WalkDirection int

const ( // order matters, (dir + 1) % 4 to rotate counter clockwise
	UP    WalkDirection = 0
	RIGHT               = 1
	DOWN                = 2
	LEFT                = 3
)

func getCellVisited(mapData [][]MapCell, guardPos [2]int) int {
	i, j := guardPos[0], guardPos[1]
	direction := UP
	cellsVisited := 1 // 1 for start position
	for {
		nextMove := getNextMove(direction) // [y, x], y = -1 -> go up, x = 1 -> go right
		nextPos := [2]int{i + nextMove[0], j + nextMove[1]}
		if !checkBundaries(nextPos, len(mapData[0]), len(mapData)) {
			break
		}
		nextCell := mapData[nextPos[0]][nextPos[1]]
		if nextCell == WALL {
			direction = getNextDirection(direction)
			continue
		} else if nextCell == EMPTY {
			mapData[nextPos[0]][nextPos[1]] = SEEN
			cellsVisited += 1
		} // else nothing
		i, j = nextPos[0], nextPos[1]
	}
	return cellsVisited
}

func getInput6() ([2]int, [][]MapCell) {
	data, _ := internal.ReadFileLines("resources/day_6")
	inputMap := make([][]MapCell, len(data))
	guardPosition := [2]int{}
	for i, line := range data {
		inputMap[i] = make([]MapCell, len(data[i]))
		for j, c := range line {
			if c == '^' {
				guardPosition[0], guardPosition[1] = i, j
				inputMap[i][j] = SEEN
			} else if c == '.' {
				inputMap[i][j] = EMPTY
			} else {
				inputMap[i][j] = WALL
			}
		}
	}
	return guardPosition, inputMap
}

func Day6_1() int {
	guardPos, mapData := getInput6()
	var nCellVisited int = getCellVisited(mapData, guardPos)
	return nCellVisited
}

func Day6_2() int {
	guardPos, mapData := getInput6()
	var nObstacles int = getNObstacles(mapData, guardPos)
	return nObstacles
}

func getNextMove(direction WalkDirection) [2]int {
	nextMove := [2]int{}
	if direction == UP {
		nextMove[0], nextMove[1] = -1, 0
	} else if direction == RIGHT {
		nextMove[0], nextMove[1] = 0, 1
	} else if direction == DOWN {
		nextMove[0], nextMove[1] = 1, 0
	} else if direction == LEFT {
		nextMove[0], nextMove[1] = 0, -1
	}
	return nextMove
}

func isWallOnPath(mapData [][]MapCell, i, j int, direction WalkDirection) bool {
	move := getNextMove(direction)
	nextPos := [2]int{i, j}
	for {
		nextPos[0], nextPos[1] = nextPos[0]+move[0], nextPos[1]+move[1]
		if !checkBundaries(nextPos, len(mapData), len(mapData[0])) {
			return false
		}
		if mapData[nextPos[0]][nextPos[1]] == WALL {
			return true
		}
	}
}

func getNextDirection(direction WalkDirection) WalkDirection {
	return (direction + 1) % 4
}

func checkBundaries(nextPos [2]int, lenWidth int, lenHeight int) bool {
	return !(nextPos[0] < 0 || nextPos[0] >= lenHeight || nextPos[1] < 0 || nextPos[1] >= lenWidth)
}

func getNObstacles(mapData [][]MapCell, guardPos [2]int) int {
	obstacles := 0
	i, j := guardPos[0], guardPos[1]
	direction := UP
	for {
		nextMove := getNextMove(direction)
		nextPos := [2]int{i + nextMove[0], j + nextMove[1]}
		if !checkBundaries(nextPos, len(mapData), len(mapData[0])) {
			break
		}
		nextCell := mapData[nextPos[0]][nextPos[1]]
		if nextCell == WALL {
			direction = getNextDirection(direction)
			continue
		} else if nextCell == EMPTY {
			mapData[nextPos[0]][nextPos[1]] = SEEN
			// if !(nextPos[0] == guardPos[0] && nextPos[1] == guardPos[1]) {
			if loopFromHere(mapData, i, j, direction, nextPos) {
				obstacles += 1
			}
			// }
		}

		i, j = nextPos[0], nextPos[1]
	}
	return obstacles
}

func loopFromHere(mapData [][]MapCell, i, j int, direction WalkDirection, obstaclePos [2]int) bool {
	// If an obstacle is put next cell, check if we hit a wall after rotation, if not, dont put an obstacle here
	nextDirection := getNextDirection(direction)
	if !isWallOnPath(mapData, i, j, nextDirection) {
		return false
	}
	// Put the obstacle
	mapData[obstaclePos[0]][obstaclePos[1]] = WALL // temporarly place a wall
	isALoop := simulateLoop(mapData, i, j, direction)
	mapData[obstaclePos[0]][obstaclePos[1]] = SEEN // reset wall
	return isALoop
}

func simulateLoop(mapData [][]MapCell, i, j int, direction WalkDirection) bool {
	cellsVisited := make(map[int]WalkDirection)
	// startPos, startDirection := [2]int{i, j}, direction
	direction = getNextDirection(direction)
	for {
		v, ok := cellsVisited[(i<<8)+j]
		if !ok {
			cellsVisited[(i<<8)+j] = direction
		} else {
			if v == direction {
				return true
			}
		}
		nextMove := getNextMove(direction)
		nextPos := [2]int{i + nextMove[0], j + nextMove[1]}
		if !checkBundaries(nextPos, len(mapData[0]), len(mapData)) {
			return false
		}
		nextCell := mapData[nextPos[0]][nextPos[1]]
		if nextCell == WALL {
			direction = getNextDirection(direction)
			continue
		} // else nothing
		i, j = nextPos[0], nextPos[1]
	}
}
