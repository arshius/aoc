package problem1

import (
	"bufio"
	"container/heap"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type PriorityQueue []*Entry

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].value < pq[j].value
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	entry := x.(*Entry)
	*pq = append(*pq, entry)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	entry := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return entry
}

type Input struct {
	left PriorityQueue
	right PriorityQueue
}

type Entry struct {
	value	int
}

type Frequency struct {
	freq map[int]int
	entry []int
	count int
}

func newInput(length int) *Input  {
	return &Input{
		make(PriorityQueue, length),
		make(PriorityQueue, length),
	}
}

func newFrequency(length int) *Frequency {
	return &Frequency{
		make(map[int]int, length),
		make([]int,0, length),
		0,
	}
}

var FileNotFound = errors.New("File not found (./1/input)")
var FileReadError = errors.New("Encounter an error when reading file")

func Solve() (int, error) {
	input, err := os.Open("./1/input")
	if err != nil {
		return 0, FileNotFound
	}
	defer input.Close()

	content := newInput(1000)
	i := 0

	reader := bufio.NewScanner(input)
	for reader.Scan() {
		line :=strings.Split(reader.Text(), "   ")
		left, _ :=	strconv.Atoi(line[0])
		right, _ := strconv.Atoi(line[1])
		content.left[i] = &Entry{
			left,
		}
		content.right[i] = &Entry{
			right,
		}
		i++
	}
	
	if err := reader.Err(); err != nil {
		return 0, FileReadError 
	}

	heap.Init(&content.left)
	heap.Init(&content.right)

	result := 0
	for content.left.Len() > 0 {
		left := heap.Pop(&content.left).(*Entry).value
		right := heap.Pop(&content.right).(*Entry).value
		delta := left - right
		if delta < 0 {
			delta = ^delta + 1 // flip all bits if negative and add 1 to make the number absolute
		}
		result += delta
	}

	return result, nil
}

func Solve2() (int, error) {
	input, err := os.Open("./1/input")
	if err != nil {
		return 0, FileNotFound
	}
	defer input.Close()

	contentL := newFrequency(9999)
	contentR := newFrequency(9999)
	reader := bufio.NewScanner(input)
	for reader.Scan() {
		line := strings.Split(reader.Text(), "   ")
		left, _:= strconv.Atoi(line[0])
		right, _:= strconv.Atoi(line[1])
		if contentL.freq[left] == 0 {
			contentL.count++
			contentL.entry = append(contentL.entry, left)
		}
		if contentR.freq[right] == 0 {
			contentR.count++
			contentR.entry = append(contentR.entry, right)
		}
		contentL.freq[left]++
		contentR.freq[right]++
	}

	if err := reader.Err(); err != nil {
		return 0, FileReadError
	}

	result := 0
	for _, value := range contentL.entry {
		if contentR.freq[value] > 0 {
			result += value * contentL.freq[value] * contentR.freq[value]
		}	
	} 

	return result, nil
}
