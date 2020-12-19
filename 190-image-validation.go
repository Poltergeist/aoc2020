package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/poltergeist/aoc2020/fileLoader"
)

func generateValid(v []string, r map[string]string) []string {
	return []string{}
}

func main() {
	input := strings.Split(string(fileLoader.LoadFile("190-input")), "\n\n")
	rules := map[string]string{}
	rx := regexp.MustCompile(`[a-zA-Z]`)
	for _, r := range strings.Split(input[0], "\n") {
		ra := strings.Split(r, ":")
		switch rx.MatchString(ra[1]) {
		case true:
			rules[ra[0]] = rx.FindString(ra[1])

		case false:
			rules[ra[0]] = strings.Trim(ra[1], " ")
		}

	}
	rx = regexp.MustCompile(`[0-9]`)
	rx = regexp.MustCompile("^" + getRulesString(rules["0"], rules, rx) + "$")

	v := 0
	for _, l := range strings.Split(input[1], "\n") {

		if rx.MatchString(l) {
			v++
		}
	}
	fmt.Println(v)
}

func getRulesString(rule string, rules map[string]string, rx *regexp.Regexp) string {
	if rx.MatchString(rule) {
		ret := []string{}
		for _, r := range strings.Split(rule, "|") {
			x := ""
			for _, rn := range strings.Split(strings.Trim(r, " "), " ") {
				x += getRulesString(rules[rn], rules, rx)

			}
			ret = append(ret, x)
		}
		return "(" + strings.Join(ret, "|") + ")"

	}

	return rule
}
