package main

import (
	"strconv"
	"strings"
)

type Stack struct {
	record []int
}

func exclusiveTime(n int, logs []string) []int {
	res := make([]int, n)
	s := &Stack{}

	prevPos := 0
	for _, log := range logs {
		record := strings.Split(log, ":")
		currID, _ := strconv.Atoi(record[0])
		currPos, _ := strconv.Atoi(record[2])
		if record[1] == "start" {
			if prevID, ok := s.Peek(); ok {
				res[prevID] += currPos - prevPos - 1
			}
			s.Push(currID)
		} else {
			runID, _ := s.Pop()
			res[runID] += currPos - prevPos + 1
		}
		prevPos = currPos
	}
	return res
}

func (s *Stack) Push(x int) {
	s.record = append(s.record, x)
}

func (s *Stack) Pop() (int, bool) {
	if len(s.record) == 0 {
		return -1, false
	}
	peek := s.record[len(s.record)-1]
	s.record = s.record[:len(s.record)-1]
	return peek, true
}

func (s *Stack) Peek() (int, bool) {
	if len(s.record) == 0 {
		return -1, false
	}
	peek := s.record[len(s.record)-1]
	return peek, true
}
