package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/poltergeist/aoc2020/fileLoader"
)

type tile struct {
	id    int
	sides map[string]string
}

func Reverse(s string) string {
	r := []rune(s)
	var output strings.Builder
	for i := len(r) - 1; i >= 0; i-- {
		output.WriteString(string(r[i]))
	}

	return output.String()
}

func main() {
	input := strings.Split(strings.Trim(string(fileLoader.LoadFile("200-input")), "\n"), "\n\n")

	tiles := make([]tile, len(input))

	for k, t := range input {
		tileId, content := func() (int, []string) {
			x := strings.Split(t, ":\n")
			i, _ := strconv.Atoi(strings.Split(x[0], " ")[1])
			return i, strings.Split(x[1], "\n")
		}()
		l := ""
		r := ""
		le := len(content) - 1
		for k, _ := range content {
			l += string(content[k][0])
			r += string(content[k][le])
		}
		tiles[k] = tile{
			id: tileId,
			sides: map[string]string{
				"top":     content[0],
				"right":   r,
				"bottom":  content[le],
				"left":    l,
				"topR":    Reverse(content[0]),
				"rightR":  Reverse(r),
				"bottomR": Reverse(content[le]),
				"leftR":   Reverse(l),
			},
		}
	}

	r := 1
	for k, t := range tiles {
		m := 0
		for kt, tt := range tiles {
			if k == kt {
				continue
			}
			for _, s := range []string{"top", "left", "bottom", "right", "topR", "leftR", "bottomR", "rightR"} {
				for _, st := range []string{"top", "left", "bottom", "right", "topR", "leftR", "bottomR", "rightR"} {
					if t.sides[s] == tt.sides[st] {
						m++
					}

				}
			}
		}
		if m == 4 {
			r *= t.id
		}
	}
	fmt.Println(r)
}
