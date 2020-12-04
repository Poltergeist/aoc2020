package main

import (
	"strings"
	"fmt"
	"github.com/poltergeist/aoc2020/fileLoader"
)

func main() {
	required := map[string]bool{"byr": true,
		"iyr":  true,
		"eyr": true,
		"hgt":  true,
		"hcl": true,
		"ecl":  true,
		"pid": true,
	}
	input := strings.Split(string(fileLoader.LoadFile("./040-input")), "\n")
	valid := 0
	rf := []string{}
	for _, l := range input {
		if l == "" {
			if len(rf) >= len(required) {
				valid = valid + 1
			}
			rf = []string{}
			continue
		}
		for _, x := range strings.Split(l, " ") {
			f := strings.Split(x, ":")[0]
			if required[f] {
				rf = append(rf, f)
			}
		}
	}
	fmt.Println(valid)
}
