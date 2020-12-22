package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/poltergeist/aoc2020/fileLoader"
)

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

	for len(o) > 0 && len(t) > 0 {

		if o[0] > t[0] {
			o = append(o[1:], []int{o[0], t[0]}...)
			t = t[1:]
			continue
		}
		t = append(t[1:], []int{t[0], o[0]}...)
		o = o[1:]
	}
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
