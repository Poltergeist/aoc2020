package main

import (
	"fmt"
	"github.com/poltergeist/aoc2020/fileLoader"
	"strings"
)

func main() {
	temp := strings.Split(strings.Trim(string(fileLoader.LoadFile("110-input")), "\n"), "\n")
	temp = strings.Split(strings.Trim(`L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL
`, "\n"), "\n")
	input := strings.Join(temp, "")
	next := input
	for {

		l := len(temp[0])

		for k, x := range input {
			o := 0
			r := []int{(l + 1) * -1, l * -1, (l - 1) * -1, -1, 1, l - 1, l, l + 1}
			if (k+1)%10 == 1 {
				r = []int{l * -1, (l - 1) * -1, 1, l, l + 1}

			}
			if (k+1)%10 == 0 {
				r = []int{(l + 1) * -1, l * -1, -1, l - 1, l}
			}
			for _, y := range r {

				if k+y >= 0 && k+y < len(input) {
					if string(input[k+y]) == "#" {
						o++
					}

				}
			}
			switch string(x) {
			case "L":
				if o == 0 {
					next = next[:k] + "#" + next[k+1:]
				}

			case "#":
				if o >= 4 {
					next = next[:k] + "L" + next[k+1:]
				}
			}

		}
		if input == next {
			fmt.Println(strings.Count(next, "#"))
			return
		}
		input = next
	}
}
