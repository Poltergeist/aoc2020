package main

import (
	"fmt"
	"github.com/poltergeist/aoc2020/day03"
	"github.com/poltergeist/aoc2020/fileLoader"
	"strings"
)

func main() {
	input := strings.Split(string(fileLoader.LoadFile("./030-input")), "\n")
	result := []string{}
	for i := range input {
		if input[i] != "" {
			result = append(result, input[i])
		}
	}
	count := 1
	l := []day03.Move{
		{R: 1, D: 1},
		{R: 3, D: 1},
		{R: 5, D: 1},
		{R: 7, D: 1},
		{R: 1, D: 2},
	}
	for _, x := range l {
		count = count * day03.TreeCount(result, x)
	}
	fmt.Println(count)
}
