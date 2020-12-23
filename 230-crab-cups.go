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

func main() {
	input := []int{}
	low := 10000000
	high := 0
	for _, v := range strings.Split("389125467", "") {
		n := func(i string) int {
			xi, _ := strconv.Atoi(i)
			return xi
		}(v)
		if n < low {
			low = n
		}
		if n > high {
			high = n
		}
		input = append(input, n)

	}

	p := 0
	for x := 1; x <= 10; x++ {
		l := input[p]
		t := append([]int{}, input[p+1:p+4]...)
		input = append(input[:p+1], input[p+4:]...)
		d := l - 1
		for Index(input, d) == -1 {
			d--
			if d <= low {
				d = high
			}

		}

		di := Index(input, d)

		input = append(input[:di+1], append(t, input[di+1:]...)...)

		fmt.Printf(`-- move %d --
label: %d
cups: %v
pick up: %v
destination: %d
`, x, l, input, t, d)
		p = Index(input, l) + 1
		if p >= len(input) {
			p = 0
		}
		if p+4 >= len(input) {
			input = append(input[3:], input[:3]...)
			p = Index(input, l) + 1
		}

	}
	fmt.Println(input)
	io := Index(input, 1)
	fmt.Println(append(input[io+1:], input[:io]...))
}
