package main

import (
	"fmt"
	"github.com/poltergeist/aoc2020/fileLoader"
	"regexp"
	"strconv"
	"strings"
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
	input := strings.Split(strings.Trim(string(fileLoader.LoadFile("160-input")), "\n"), "\n")
	tickets := input[Index(input, "nearby tickets")+1:]
	myTicket := strings.Split(input[Index(input, "your ticket")+1], ",")

	rules := input[:Index(input, "your ticket")-1]
	type rule struct {
		l int
		h int
		n string
	}
	ruleset := []rule{}
	re, _ := regexp.Compile(`([\w\ ]*)(?:\: )(\d*)-(\d*) or (\d*)-(\d*)`)
	for _, r := range rules {
		x := re.FindStringSubmatch(r)
		xi := []int{}
		for _, v := range x[2:] {
			a, _ := strconv.Atoi(v)
			xi = append(xi, a)
		}
		ruleset = append(ruleset, []rule{rule{l: xi[0], h: xi[1], n: x[1]}, rule{l: xi[2], h: xi[3], n: x[1]}}...)
	}

	candidates := map[int]map[int]map[string]struct{}{}

	for row, t := range tickets {
		for k, tf := range strings.Split(t, ",") {
			if _, ok := candidates[k]; !ok {
				candidates[k] = map[int]map[string]struct{}{}
			}
			if _, ok := candidates[k][row]; !ok {
				candidates[k][row] = map[string]struct{}{}
			}
			v, _ := strconv.Atoi(tf)
			valid := false
			pf := map[string]struct{}{}
			for _, r := range ruleset {
				if v >= r.l && v <= r.h {
					pf[r.n] = struct{}{}
					valid = true
				}
			}
			if valid {
				candidates[k][row]=pf

			}
		}
	}

	df := map[string]int{}
	c := 0
	for len(df) != len(myTicket) && c <= 100 {
		c++
		for k, v := range candidates {
			vf := v[0]
			for c, _ := range vf {
				if _, ok := df[c]; ok {
					delete(vf, c)
				}
			}
			for row, _ := range v {
				if len(candidates[k][row]) == 0 {
					continue
				}
				for c, _ := range vf {
					if _, ok := candidates[k][row][c]; !ok {
						delete(vf, c)
					}
				}

			}
			if len(vf) == 1 {
				for x, _ := range vf {
					df[x] = k
				}
			}
		}

	}

	r := 1
	for f, i := range df{
		if strings.Contains(f, "departure") {
			x, _ := strconv.Atoi(myTicket[i])
			r *= x
		}
	}
	fmt.Println(r)
}
