/*
https://adventofcode.com/2024/day/8
*/
package solutions

import (
	"aoc/internal"
)

func getInput9() []int {
	data, _ := internal.ReadFileLines("resources/day_9")
	input := make([]int, len(data[0]))
	for i := range data[0] {
		input[i] = int(data[0][i]) - 48
	}
	return input
}

func expand(data []int) []int {
	fileExpanded := make([]int, 0, len(data)*4)
	spaceTurn := false
	fileId := 0
	for _, v := range data {
		if spaceTurn {
			for i := 0; i < v; i++ {
				fileExpanded = append(fileExpanded, -1)
			}
		} else {
			for i := 0; i < v; i++ {
				fileExpanded = append(fileExpanded, fileId)
			}
			fileId++
		}
		spaceTurn = !spaceTurn // alternate between block file and free space
	}
	return fileExpanded
}

func findLeftMostEmpty(data []int, start int) int {
	for i := start; i < len(data); i++ {
		if data[i] == -1 {
			return i
		}
	}
	return 0 // never happens
}

func findRightMostBlock(data []int, start int) int {
	for i := start; i >= 0; i-- {
		if data[i] != -1 {
			return i
		}
	}
	return 0 //never happens
}

func compactFrag(data []int) {
	leftMost := 0
	rightMost := len(data) - 1
	for {
		leftMostEmpty := findLeftMostEmpty(data, leftMost)
		rightMostBlock := findRightMostBlock(data, rightMost)
		if leftMostEmpty >= rightMostBlock {
			break
		}
		data[leftMostEmpty], data[rightMostBlock] = data[rightMostBlock], -1
	}
}

func checksum(data []int) int {
	total := 0
	for i := 1; i < len(data); i++ {
		if data[i] == -1 {
			continue
		}
		total += i * data[i]
	}
	return total
}

func solve9(data []int, part2 bool) int {
	fileExpanded := expand(data)
	if part2 {
		compactFile(fileExpanded)

	} else {
		compactFrag(fileExpanded)
	}
	checksum := checksum(fileExpanded)
	return checksum
}

func Day9_1() int {
	data := getInput9()
	soluce := solve9(data, false)
	return soluce
}

func findEmptyBlock(data []int, size int, end int) int {
	blockSize := 0
	blockStart := 0
	for i := 0; i < end; i++ {
		if data[i] != -1 { // skip files
			blockSize = 0
			blockStart = i
			continue
		}
		blockSize += 1
		if blockSize == size {
			return blockStart + 1
		}
	}
	return -1
}

func findNextFile(data []int, start, blockIdToFind int) (int, int, int) {
	if start == 1 { // the end
		return 0, 0, 0
	}
	i := start - 1
	for ; i >= 0; i-- { // skip blanks
		if data[i] == -1 || data[i] != blockIdToFind {
			continue
		} else {
			break
		}
	}
	blockId := data[i]
	if i <= 0 || blockId == 0 {
		return 0, 0, 0
	}
	endBlock := i
	for ; i >= 0; i-- {
		if data[i] != blockId {
			break
		}
	}
	return i + 1, endBlock - i, blockId
}

func compactFile(data []int) {
	blockPos, blockId := len(data), data[len(data)-1]+1
	var blockSize int
	for {
		blockPos, blockSize, blockId = findNextFile(data, blockPos, blockId-1)
		if blockPos == 0 {
			break
		}
		leftMostEmpty := findEmptyBlock(data, blockSize, blockPos)
		if leftMostEmpty != -1 {
			for i := 0; i < blockSize; i++ {
				data[leftMostEmpty+i], data[blockPos+i] = blockId, -1
			}
		}
	}
	// fmt.Println(data)
}

func Day9_2() int {
	data := getInput9()
	soluce := solve9(data, true)
	return soluce
}
