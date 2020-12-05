package main

import (
	"fmt"
	"strings"
	"github.com/poltergeist/aoc2020/fileLoader"
	"math"
	"sort"
)
func Index(vs []int, t int) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}
func Include(vs []int, t int) bool {
	return Index(vs, t) >= 0
}
func main() {
	input := strings.Split(string(fileLoader.LoadFile("./050-input")), "\n")
	// input := strings.Split(`BBFFBBFRLL`, "\n")
	sl := []int {}

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
		sl = append(sl, int(id))
	}
	sort.Ints(sl)
	for _, i := range sl {
		if !Include(sl, i+1) && Include(sl, i+2) {
			fmt.Println("my ID",  i+1)
		}
	}
}
