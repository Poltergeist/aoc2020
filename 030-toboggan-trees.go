package main

import (
	"fmt"
	"github.com/poltergeist/aoc2020/day03"
	"strings"
	"github.com/poltergeist/aoc2020/fileLoader"
)

func main() {
	input := strings.Split(string(fileLoader.LoadFile("./030-input")), "\n")
	result := []string{}
	for i := range input {
		if input[i] != "" {
			result  = append(result, input[i])
		}
	}
	fmt.Println(day03.TreeCount(result, day03.Move{R: 3, D:1}))
}
