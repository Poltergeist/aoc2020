package main

import (
	"fmt"
	"strconv"
	"strings"
	"github.com/poltergeist/aoc2020/fileLoader"
)

func main() {
	acc := 0
	input := strings.Split(strings.Trim(string(fileLoader.LoadFile("./080-input")),"\n"), "\n")
	p := 0
	ch := map[int]struct{}{}
	for p < len(input)  {
		c := strings.Split(input[p], " ")
		if _, ok := ch[p]; ok {
			break
		}
		ch[p] = struct{}{}
		if c[0] == "acc" {
			i, _ := strconv.Atoi(c[1])
			acc += i
		}
		if c[0] == "jmp" {
			i, _ := strconv.Atoi(c[1])
			p += i
			continue
		}
		p++
	}
	fmt.Println(acc)
}
