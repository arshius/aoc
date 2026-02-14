package problem3

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func Part1(input *os.File) int {
	//text := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	var regex = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	solution := 0

	reader := bufio.NewScanner(input)
	for reader.Scan() {
		text := reader.Text()
		matches := regex.FindAllStringSubmatch(text, -1)
		for _, match := range matches {
			a, _ := strconv.Atoi(match[1])
			b, _ := strconv.Atoi(match[2])
			solution += a * b
		}
	}
	return solution
}

func Part2(input *os.File) int {
	//text := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
	var regex = regexp.MustCompile(`(?:mul\((\d{1,3}),(\d{1,3})\))|(do(?:n't)?\(\))`)
	solution := 0
	disabled := false

	reader := bufio.NewScanner(input)
	for reader.Scan() {
		text := reader.Text()
		matches := regex.FindAllStringSubmatch(text, -1)
		for _, match := range matches {
			fmt.Printf("%v\n", match)

			if match[3] == "do()" {
				disabled = false
				fmt.Println("resume")
				continue
			} else if match[3] == "don't()" {
				disabled = true
				fmt.Println("pause")
				continue
			}

			if !disabled {
				a, _ := strconv.Atoi(match[1])
				b, _ := strconv.Atoi(match[2])
				solution += a * b
				fmt.Printf("%d * %d = %d\n", a, b, solution)
			}
		}
	}

	return solution
}

func Solve() (int, error) {
	input, error := os.Open("./3/input")
	if error != nil {
		return 0, errors.New("File not found ./3/input")
	}
	defer input.Close()

	solution := Part2(input)

	return solution, nil
}
