package main

import (
	"strconv"
	"strings"
)

type log struct {
	num    int
	action string
	pos    int
}

func exclusiveTime(n int, logs []string) []int {
	records := Record(logs)
	res := make([]int, n)
	prev := 0
	stack := []log{records[0]}
	for i := 1; i < len(records); i++ {
		if records[i].action == "start" {
			if len(stack) != 0 {
				res[stack[len(stack)-1].num] += records[i].pos - prev
			}
			stack = append(stack, records[i])
			prev = records[i].pos
		} else {
			res[stack[len(stack)-1].num] += records[i].pos - prev + 1
			stack = stack[:len(stack)-1]
			prev = records[i].pos + 1
		}
	}
	return res
}

func Record(logs []string) []log {
	var records []log
	for _, l := range logs {
		s := strings.Split(l, ":")
		def, _ := strconv.Atoi(s[0])
		pos, _ := strconv.Atoi(s[2])
		records = append(records, log{def, s[1], pos})
	}
	return records
}
