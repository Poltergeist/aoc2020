package main

import (
	"fmt"
	"github.com/poltergeist/aoc2020/fileLoader"
	"strings"
)

func main() {
	input := strings.Split(strings.Trim(string(fileLoader.LoadFile("./060-input")), "\n"), "\n\n")
	sum := 0
	for _, x := range input {
		first := ""
		for k, a := range strings.Split(x, "\n") {
			if k == 0 {
				first = a
				continue
			}
			for _, c := range first {
				if !strings.Contains(a, string(c)) {
					first = strings.ReplaceAll(first, string(c), "")
				}
			}
		}
		sum = sum + len(first)
	}
	fmt.Println(sum)
}
