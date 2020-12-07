package main

import (
	"fmt"
	"regexp"
	"strings"
	"github.com/poltergeist/aoc2020/fileLoader"
)

func main() {
	input := string(regexp.MustCompile(`([,.\d\ ])|(bags|bag)`).ReplaceAll([]byte(strings.Trim(string(fileLoader.LoadFile("./070-input")),"\n")), []byte("")))

	canContain := map[string]struct{}{"shinygold":struct{}{}}

	b := 0
	for b != len(canContain)  {
		b = len(canContain)

		for _, l := range strings.Split(input, "\n") {
			r := strings.Split(l, "contain")
			for k, _ := range canContain {
				if strings.Contains(r[1], k) {
					canContain[r[0]] = struct{}{}
				}

			}
		}
	}
	delete(canContain, "shinygold")
	fmt.Println(len(canContain))
}
