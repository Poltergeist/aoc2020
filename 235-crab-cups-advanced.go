package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Index(vs []int, t int) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

func toInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

type cupNode struct {
	val   int
	right *cupNode
}

func main() {

	input := strings.Split("942387615", "")
	for len(input) < 1000000 {
		input = append(input, fmt.Sprintf("%d", len(input)+1))
	}
	low := 1
	high := len(input)
	p := &cupNode{
		val: toInt(input[0]),
	}
	prev := p
	cache := map[int]*cupNode{p.val: p}
	for _, v := range input[1:] {
		prev.right = &cupNode{
			val: toInt(v),
		}

		prev = prev.right
		cache[prev.val] = prev

	}
	prev.right = p

	for x := 0; x < 10000000; x++ {
		temp := p.right
		p.right = temp.right.right.right

		d := p.val - 1
		for d == 0 || temp.val == d || temp.right.val == d || temp.right.right.val == d {
			d--
			if d < low {
				d = high
			}

		}
		temp.right.right.right = cache[d].right
		cache[d].right = temp

		p = p.right
	}

	fmt.Println(cache[1].right.val * cache[1].right.right.val)
}
