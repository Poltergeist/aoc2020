package main

import (
	"fmt"
	"strings"
	"github.com/poltergeist/aoc2020/fileLoader"
	"math"
)

func main() {
	input := strings.Split(string(fileLoader.LoadFile("./050-input")), "\n")
	// input := strings.Split(`BBFFBBFRLL`, "\n")

	max := 0.0
	for _, s := range input {
		if s == "" {
			continue
		}
		lower := 0.0
		higher := 127.0
		row := 0.0
		for _, r := range s[:7] {
			if string(r) == "B" {
				lower = lower + 1 + math.Floor((higher-lower)/2)
				row = higher
				continue
			}
			higher = higher - math.Ceil((higher-lower)/2)
			row = lower
		}
		lower = 0.0
		higher = 7.0
		seat := 0.0
		for _, r := range s[7:] {
			if string(r) == "R" {
				lower = lower + 1 + math.Floor((higher-lower)/2)
				seat = higher
				continue
			}
			higher = higher - math.Ceil((higher-lower)/2)
			seat = lower
		}
		id := row*8 + seat
		if id > max {
			max = id
		}
	}
	fmt.Println(max)
}
