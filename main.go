package main

import (
	"aoc/solutions"
	"fmt"
	"os"
)

func main() {
	ret := run(os.Args)
	os.Exit(ret)
}

func run(args []string) int {
	solution := solutions.Day8_1()
	fmt.Printf("1st part: %d\n", solution)
	solution = solutions.Day8_2()
	fmt.Printf("2nd part: %d\n", solution)

	return 0
}
