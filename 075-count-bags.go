package main

import (
	"fmt"
	"github.com/poltergeist/aoc2020/fileLoader"
	"regexp"
	"strconv"
	"strings"
)

func getBagCount(l map[string]string, bn string, bc int) int {
	contents, ok := l[bn]
	if !ok {
		return bc
	}
	total := bc
	for _, v := range strings.Split(contents, ",") {
		c, _ := strconv.Atoi(string(regexp.MustCompile(`\D`).ReplaceAll([]byte(v), []byte(""))))
		n := string(regexp.MustCompile(`\d`).ReplaceAll([]byte(v), []byte("")))
		total += getBagCount(l, n, c) * bc
	}
	return total
}

func main() {
	input := string(regexp.MustCompile(`([.\ ])|(bags|bag)`).ReplaceAll([]byte(strings.Trim(string(fileLoader.LoadFile("./070-input")), "\n")), []byte("")))

	canContain := map[string]string{}

	for _, l := range strings.Split(input, "\n") {
		r := strings.Split(l, "contain")
		if r[1] != "noother" {
			canContain[r[0]] = r[1]
		}
	}
	fmt.Println(getBagCount(canContain, "shinygold", 1) - 1)
}
