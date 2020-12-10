package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"github.com/poltergeist/aoc2020/fileLoader"
)

func main() {
	temp := strings.Split(strings.Trim(string(fileLoader.LoadFile("100-input")), "\n"), "\n")
// 	temp := strings.Split(strings.Trim(`28
// 33
// 18
// 42
// 31
// 14
// 46
// 20
// 48
// 47
// 24
// 23
// 49
// 45
// 19
// 38
// 39
// 11
// 1
// 32
// 25
// 35
// 8
// 17
// 7
// 9
// 4
// 2
// 34
// 10
// 3
// `, "\n"), "\n")
	input := []int{}
	for _, i := range temp {
		x, _ := strconv.Atoi(i)
		input = append(input, x)
	}
	sort.Ints(input)
	diff := map[int]int{1: 0, 2: 0, 3: 1}

	b := 0

	for _, x := range input {
		d := x-b
		b = x
		diff[d]++
	}


	fmt.Println(input, diff, diff[1] * diff[3])
}
