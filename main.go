package main

import (
	"aoc/solutions"
	"fmt"
	"os"
	"time"
)

func main() {
	ret := run(os.Args)
	os.Exit(ret)
}

func run(args []string) int {
	start := time.Now()
	solution := solutions.Day15_1()
	elapsed := time.Since(start)
	fmt.Printf("1st part: %d in %dus\n", solution, elapsed.Microseconds())
	start = time.Now()
	solution = solutions.Day15_2()
	elapsed = time.Since(start)
	fmt.Printf("2nd part: %d in %dus\n", solution, elapsed.Microseconds())
	return 0
}
