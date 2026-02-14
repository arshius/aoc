package main

import (
	"fmt"

	"aoc/aoc2024/3"
)

func main() {
	solve, err := problem3.Solve()
	if err != nil {
		fmt.Printf("%v", err)
	}

	fmt.Printf("solution is %d\n",solve)
}
