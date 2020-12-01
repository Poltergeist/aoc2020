package main

import (
	"fmt"
	"github.com/poltergeist/aoc2020/fileLoader"
	"strings"
	"github.com/poltergeist/aoc2020/day01"
	"github.com/poltergeist/aoc2020/misc"
)

func main(){
	input := strings.Split(string(fileLoader.LoadFile("./010-input")), "\n")
	fmt.Println(day01.Checksum(2020, misc.ConvertStringsToInt(input)))
}
