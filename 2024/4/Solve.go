package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func MockSolve1() int {
	mockText :=
		`MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

	input := strings.NewReader(mockText)
	reader := bufio.NewScanner(input)

	origin, vert, ldia, rdia := read(reader)
	full := append(append(append(origin, vert...), ldia...), rdia...)
	return scan(full)
}

func MockSolve2() int {
	mockText :=
		`MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

	input := strings.NewReader(mockText)
	reader := bufio.NewScanner(input)

	origin, _, ldia, rdia := read(reader)
	return scanX(rdia, ldia, len(origin[0]))
}

func Solve1() int {
	input, error := os.Open("./input")
	if error != nil {
		fmt.Println("File not found ./input")
	}
	defer input.Close()

	reader := bufio.NewScanner(input)

	origin, vert, ldia, rdia := read(reader)
	full := append(append(append(origin, vert...), ldia...), rdia...)
	return scan(full)
}

func Solve2() int {
	input, error := os.Open("./input")
	if error != nil {
		fmt.Println("File not found ./input")
	}
	defer input.Close()

	reader := bufio.NewScanner(input)

	origin, _, ldia, rdia := read(reader)
	return scanX(rdia, ldia, len(origin[0]))
}

func read(reader *bufio.Scanner) ([]string, []string, []string, []string) {
	var origin, vert, ldia, rdia []string

	i := 0
	for reader.Scan() {
		data := reader.Text()
		width := len(data) 
		if origin == nil {
			origin = make([]string, 0)
			vert = make([]string, width)
			ldia = make([]string, width*2-1)
			rdia = make([]string, width*2-1)
		}
		origin = append(origin, data)
		transform(i, width-1, data, &vert, &ldia, &rdia)
		i++
	}

	//fmt.Printf("%v\n", vert)
	//fmt.Printf("%v\n", rdia)
	//fmt.Printf("%v\n", ldia)

	return origin, vert, ldia, rdia
}

type XMAS struct {
	n     int
	line  string
	index [][]int
}

func scanX(rdia []string, ldia []string, width int) int {
	mas := regexp.MustCompile(`(?:M(A)S)`)
	sam := regexp.MustCompile(`(?:S(A)M)`)
	total := 0

	center := XMAS{}
	for i, line := range rdia {
		index1 := mas.FindAllStringSubmatchIndex(line, -1)
		index2 := sam.FindAllStringSubmatchIndex(line, -1)
		center.n = i
		center.line = line
		center.index = append(center.index, index1...)
		center.index = append(center.index, index2...)
		factor := width - 1

		for _, ri := range center.index {
			rx := ri[0] + 1
			rn := center.n
			fmt.Printf("--- %d ---\n[%d] > %s\n%v\n", width, rn, center.line, ri)
			pos := rn + ri[2]*(width+1)
			if rn-width >= 0 {
				rn -= width
				pos += rn*factor
				rn += 1
			}

			ln := rx*2 + rn
			line = ldia[ln]
			index1 = mas.FindAllStringSubmatchIndex(line, -1)
			index2 = sam.FindAllStringSubmatchIndex(line, -1)

			fmt.Printf("pair [%d]\n[%d] > %s\n%v\n", rx, ln, line, append(index1, index2...))
			for _, li := range append(index1, index2...) {
				pos2 := ln + li[2]*factor
				if ln-width >= 0 {
					fmt.Printf("%d ", pos2)
					pos2 += (ln - factor) * factor
					fmt.Println(pos2)
				}

				fmt.Printf("%v %d %d\n", li, pos, pos2)
				if pos == pos2 {
					total++
					break
				}
			}

		}

		center.index = nil
	}

	return total
}

func scan(data []string) int {
	xmas := regexp.MustCompile(`XMAS`)
	samx := regexp.MustCompile(`SAMX`)

	total := 0

	for _, line := range data {
		fmt.Printf("----\n%s\n", line)
		matches := xmas.FindAllStringSubmatch(line, -1)
		indexs := xmas.FindAllStringSubmatchIndex(line, -1)
		fmt.Printf("%v %v\n", matches, indexs)
		total += len(matches)
		matches = samx.FindAllStringSubmatch(line, -1)
		indexs = samx.FindAllStringSubmatchIndex(line, -1)
		fmt.Printf("%v %v\n", matches, indexs)
		total += len(matches)
	}

	return total
}

func transform(line int, length int, origin string, vert *[]string, ldia *[]string, rdia *[]string) {
	for n, c := range origin {
		(*vert)[n] += string(c)
		if n-line < 0 {
			(*rdia)[length+line-n] += string(c)
		} else {
			(*rdia)[n-line] += string(c)
		}
		(*ldia)[n+line] += string(c)
		//fmt.Printf("%c\n", line)
	}
}

func main() {
	solution := Solve2()
	fmt.Printf("Solution is : %d\n", solution)
}
