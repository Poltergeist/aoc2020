package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/poltergeist/aoc2020/fileLoader"
)

func play(o []int, t []int, d int) ([]int, []int) {
	co := map[string]bool{}
	ct := map[string]bool{}
	for len(o) > 0 && len(t) > 0 {
		if co[fmt.Sprintf("%v", o)] || ct[fmt.Sprintf("%v", t)] {
			o = append(o[1:], []int{o[0], t[0]}...)
			t = t[1:]
			continue
		}
		co[fmt.Sprintf("%v", o)] = true
		ct[fmt.Sprintf("%v", t)] = true

		if o[0] <= len(o)-1 && t[0] <= len(t)-1 {
			to, _ := play(append([]int{}, o[1:o[0]+1]...), append([]int{}, t[1:t[0]+1]...), d+1)
			if len(to) != 0 {
				o = append(o[1:], []int{o[0], t[0]}...)
				t = t[1:]
				continue
			}
			t = append(t[1:], []int{t[0], o[0]}...)
			o = o[1:]
			continue

		}

		if o[0] > t[0] {
			o = append(o[1:], []int{o[0], t[0]}...)
			t = t[1:]
			continue
		}
		t = append(t[1:], []int{t[0], o[0]}...)
		o = o[1:]
	}
	return o, t
}

func main() {
	input := strings.Split(strings.Trim(string(fileLoader.LoadFile("220-input")), "\n"), "\n\n")
	o := []int{}
	for _, x := range strings.Split(strings.Split(input[0], ":\n")[1], "\n") {
		xi, _ := strconv.Atoi(x)
		o = append(o, xi)
	}
	t := []int{}
	for _, x := range strings.Split(strings.Split(input[1], ":\n")[1], "\n") {
		xi, _ := strconv.Atoi(x)
		t = append(t, xi)
	}

	o, t = play(o, t, 0)

	r := 0
	if len(o) > 0 {
		for k, v := range o {
			r += (len(o) - k) * v
		}
	} else {

		for k, v := range t {
			r += (len(t) - k) * v
		}
	}

	fmt.Println(r)
}
