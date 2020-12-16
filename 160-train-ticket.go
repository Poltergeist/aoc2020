package main

import (
	"fmt"
	"regexp"
	"strings"
	"strconv"
	"github.com/poltergeist/aoc2020/fileLoader"
)

func Index(vs []string, t string) int {
	for i, v := range vs {
		if strings.Contains(v, t) {
			return i
		}
	}
	return -1
}

func main() {
	input:= strings.Split(strings.Trim(string(fileLoader.LoadFile("160-input")), "\n"), "\n")
// 	input = strings.Split(strings.Trim(`class: 1-3 or 5-7
// row: 6-11 or 33-44
// seat: 13-40 or 45-50

// your ticket:
// 7,1,14

// nearby tickets:
// 7,3,47
// 40,4,50
// 55,2,20
// 38,6,12
// `, "\n"), "\n")
	tickets := input[Index(input, "nearby tickets")+1:]
	rules := input[:Index(input, "your ticket")-1]
	type rule struct {
		l int
		h int
	}
	ruleset := []rule{}
	re, _ := regexp.Compile(`(?:\w\: )(\d*)-(\d*) or (\d*)-(\d*)`)
	for _, r := range rules {
		x := re.FindStringSubmatch(r)
		xi := []int{}
		for _,v := range x[1:] {
			a, _ := strconv.Atoi(v)
			xi = append(xi, a)
		}
		ruleset = append(ruleset, []rule{rule{l: xi[0], h: xi[1]}, rule{l: xi[2], h: xi[3]}}...)
	}

	iv := 0

	for _, t := range tickets {
		for _, tf := range strings.Split(t, ",") {
			v, _ := strconv.Atoi(tf)
			valid:= false
			for _, r := range ruleset{
				if v >= r.l && v <= r.h {
					valid = true
					break
				}
			}
			if !valid {
				iv += v
			}
		}
	}

	fmt.Println(iv)
}
