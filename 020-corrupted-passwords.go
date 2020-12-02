package main

import (
	"fmt"
	"github.com/poltergeist/aoc2020/day02"
	"strings"
	"github.com/poltergeist/aoc2020/fileLoader"
)

func main() {
	input := strings.Split(string(fileLoader.LoadFile("./020-input")), "\n")
	fmt.Println(day02.Check(input))
}
