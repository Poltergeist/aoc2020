package main

import (
	"fmt"
	"github.com/poltergeist/aoc2020/fileLoader"
	"regexp"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func move(d int, x int, y int, dis int) (int, int) {
	switch d {
	case 0:
		y += dis
		return x, y
	case 1:
		x += dis
		return x, y
	case 2:
		y -= dis
		return x, y
	case 3:
		x -= dis
		return x, y
	}
	return 0, 0
}

func main() {
	compass := "NESW"
	input := strings.Split(strings.Trim(string(fileLoader.LoadFile("120-input")), "\n"), "\n")
	s := struct {
		x int
		y int
	}{x: 0, y: 0}
	w := struct {
		x int
		y int
	}{x: 10, y: 1}

	re, _ := regexp.Compile(`([NSEWLRF])([-\d]*)`)
	for _, r := range input {
		m := re.FindStringSubmatch(r)
		a, _ := strconv.Atoi(m[2])
		c := m[1]

		if strings.Contains(compass, c) {
			w.x, w.y = move(strings.Index(compass, c), w.x, w.y, a)
		}

		if c == "F" {
			s.x += w.x * a
			s.y += w.y * a
		}

		if strings.Contains("LR", c) {
			if a == 180 {
				w.x = -w.x
				w.y = -w.y
				continue
			}
			mul := 1
			if ((c == "L") && (a == 90)) || ((c == "R") && (a == 270)) {
				mul = -1
			}
			t := w.x
			w.x = w.y * mul
			w.y = t * -mul
		}

	}

	fmt.Println(s, abs(s.x)+abs(s.y))

}
