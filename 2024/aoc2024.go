package main

import (
	"fmt"

	"aoc/aoc2024/problem1"
)

func main() {
	solve, err := problem1.Solve2()
	if err != nil {
		fmt.Printf("%v", err)
	}

	fmt.Printf("solution is %d\n",solve)
}
