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
	cha := map[int]struct{}{}
	change := false
	for p < len(input)  {
		c := strings.Split(input[p], " ")
		if _, ok := ch[p]; ok {
			p = 0
			acc = 0
			change = false
			ch = map[int]struct{}{}
			continue
		}
		ch[p] = struct{}{}
		if c[0] == "acc" {
			i, _ := strconv.Atoi(c[1])
			acc += i
		}
		if _, ok := cha[p]; !ok && !change {
			if c[0] == "jmp" {
				c[0] = "nop"
			}
			c[0] = "nop"
			cha[p] = struct{}{}
			change = true
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
