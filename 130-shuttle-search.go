package main

import (
	"fmt"
	"strconv"
	"strings"
	"github.com/poltergeist/aoc2020/fileLoader"
)

func main() {
	input := strings.Split(strings.Trim(string(fileLoader.LoadFile("130-input")), "\n"), "\n")
	input = strings.Split(strings.Trim(`939
7,13,x,x,59,x,31,19`, "\n"), "\n")
	now, _ := strconv.Atoi(input[0])
bus:=0
	min := int(^uint(0)  >> 1)
	for _, x := range strings.Split(input[1], ",") {
		if x == "x" {
			continue
		}
		t, _ := strconv.Atoi(x)
		tc := t
		for t <= now {
			t += tc
		}
		// fmt.Println(x, t, tc)
		if t < min {
			bus = tc
			min = t
		}

	}
	fmt.Println(min, now,  (min - now) * bus)
}
