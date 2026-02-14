package problem2

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type SortedQueue struct {
	data   []int
	compare func(int, int) bool
}

func less(a, b int) bool {
	return a < b && b < a+4
}

func more(a, b int) bool {
	return a > b && b > a-4
}

func (sq *SortedQueue) Sort()(bool) {
	var a, b int
	sorted := true
	for i:=0; i<len(sq.data)-1; i++ {
		a = sq.data[i]
		b = sq.data[i+1]
		if sq.compare == nil {
			if less(a, b) { 
				sq.compare = less
			} else if more(a, b) {
				sq.compare = more
			} else {
				sorted = false
				break
			}
		} else {
			if !sq.compare(a, b) {
				sorted = false
				break
			}
		}
	}
	return sorted 
}

func Solve() (int, error) {
	input, err := os.Open("./2/input")
	if err != nil {
		return 0, errors.New("File not found (./2/input)")
	}
	defer input.Close()

	//	input := []string{"40 42 44 47 49 50 58 65 67 70 71 72 74 75 76 77 78 81 83 85 87 91 93 96 98 99",
	//		"40 42 44 47 49 50 48 65 67 70 71 72 75 75 74 76 78 81 83 85 87 91 73 76 79 81",
	//		"99 98 96 91 84 77 69 60 58 51 50 49 48 47 39 32 30 25 24 23 22 21 20 19 15 14"}

	safe := make([]SortedQueue, 0, 1000)
	unsafe := make([]SortedQueue, 0, 1000)

	reader := bufio.NewScanner(input)
	n := 0
	for reader.Scan() {
		n++
		entry := SortedQueue{}
		for x := range strings.SplitSeq(reader.Text(), " ") {
			number, _ := strconv.Atoi(x)
			entry.data = append(entry.data, number)
		}
		valid := entry.Sort()
		if valid {
			safe = append(safe, entry)
		} else {
			ommited := SortedQueue{}
			for i:=0; i<len(entry.data); i++ {
				ommited.data = make([]int, 0)
				ommited.data = append(append(ommited.data, entry.data[:i]...), entry.data[i+1:]...)
				ommited.compare = nil
				valid = ommited.Sort()
				if valid {
					safe = append(safe, ommited)
					break
				}
				if i == len(entry.data)-1 {
					fmt.Printf("%d > %v [%d] trying\n", n, entry.data, len(entry.data))
					fmt.Printf("    > %v failed\n", ommited.data)
				}
			}
			if !valid {
				unsafe = append(unsafe, entry)
			}
		}
	}
	fmt.Printf("unsafe %d \n",  len(unsafe))
	//for n, data := range unsafe {
		//if len(data.data) > 0 {
		//	fmt.Printf("%d > %v \n", n+1, data.data)
		//}
	//}
	return len(safe), nil
}
