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
	// solution := solutions.Day1_1()
	// fmt.Printf("Day1.1: %d\n", solution)
	// solution = solutions.Day1_2()
	// fmt.Printf("Day1.2: %d\n", solution)
	solution := solutions.Day2_1()
	fmt.Printf("Day2.1: %d\n", solution)
	solution = solutions.Day2_2()
	fmt.Printf("Day2.2: %d\n", solution)
	return 0
}
