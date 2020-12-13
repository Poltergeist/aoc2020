package main

import (
	"fmt"
	"github.com/poltergeist/aoc2020/fileLoader"
	"strconv"
	"strings"
)

func main() {
	input := strings.Split(strings.Trim(string(fileLoader.LoadFile("130-input")), "\n"), "\n")
	inum := strings.Split(input[1], ",")
	t, _ := strconv.Atoi(inum[0])
	for {
		tnm := 1
		for k, x := range inum {
			if x == "x" {
				continue
			}
			xi, _ := strconv.Atoi(x)
			if (t+k)%xi != 0 {
				break
			}
			if k == len(inum)-1 {
				fmt.Println("yes", t)
				return
			}
			tnm *= xi

		}
		t += tnm
	}
}
