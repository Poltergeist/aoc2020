package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"github.com/poltergeist/aoc2020/fileLoader"
)

func abs(x int) int {
	if x <0 {
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
		f int
	}{x: 0, y: 0, f: strings.Index(compass, "E")}

	re, _ := regexp.Compile(`([NSEWLRF])([-\d]*)`)
	for _, r := range input {
		m := re.FindStringSubmatch(r)
		a, _ := strconv.Atoi(m[2])
		c := m[1]

		if strings.Contains(compass, c) {
			s.x, s.y = move(strings.Index(compass, c), s.x, s.y, a)
		}

		if c == "F" {
			s.x, s.y = move(s.f, s.x, s.y, a)
		}

		if c == "L" {
			an := s.f - (a / 90)
			an += 4
			s.f = an % len(compass)
		}
		if c == "R" {
			an := s.f + (a / 90)
			s.f = an % len(compass)
		}
	}

	fmt.Println(s, abs(s.x)+abs(s.y))

}
