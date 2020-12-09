package main

import (
	"fmt"
	"github.com/poltergeist/aoc2020/fileLoader"
	"sort"
	"strconv"
	"strings"
)

func main() {
	i := strings.Split(strings.Trim(string(fileLoader.LoadFile("./090-input")), "\n"), "\n")
	input := []int{}
	for _, x := range i {
		y, _ := strconv.Atoi(x)
		input = append(input, y)
	}
	pre := 25
	ki := 0
	inv := 0
	for k, v := range input[pre:] {
		if !check(input[k:pre+k], v) {
			ki = k + pre
			fmt.Println(v)
			inv = v
			break
		}
	}
	for k, _ := range input[:ki] {
		for kb, _ := range input[:ki] {
			if kb <= k {
				continue
			}
			if sum(input[k:kb]) == inv {
				s := input[k:kb]
				sort.Ints(s)
				fmt.Println(s[0], s[len(s)-1], s[0]+s[len(s)-1])
				return
			}

		}

	}
}
func sum(array []int) int {
	sum := 0
	for _, v := range array {
		sum += v
	}
	return sum
}

func check(l []int, needle int) bool {
	for ka, a := range l {
		for kb, b := range l {
			if ka == kb {
				continue
			}
			if a+b == needle {
				return true
			}

		}
	}
	return false
}
