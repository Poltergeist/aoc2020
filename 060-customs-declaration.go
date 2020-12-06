package main

import (
	"github.com/poltergeist/aoc2020/fileLoader"
	"fmt"
	"strings"
)

func main() {
	input := strings.Split(string(fileLoader.LoadFile("./060-input")), "\n\n")
	sum := 0
	for _, x := range input {
		group := ""
		for _, a := range strings.Split(x, "\n") {
			for _, b := range a {
				if !strings.Contains(group, string(b)) {
					group = group + string(b)
				}
			}
		}
		sum = sum + len(group)
	}
	fmt.Println(sum)
}
