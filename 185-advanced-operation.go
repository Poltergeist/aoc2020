package main

import (
	"fmt"
	"github.com/poltergeist/aoc2020/fileLoader"
	"regexp"
	"strconv"
	"strings"
)

func calc(s string) string {
	re := regexp.MustCompile(`(\d+ \+ \d+)`)

	for m := re.FindAllStringSubmatch(s, 2); len(m) >= 1 && len(m[0][0]) != len(s); m = re.FindAllStringSubmatch(s, 2) {
		rx := regexp.MustCompile(`(^|[ \D])` + regexp.MustCompile(`(\*|\+)`).ReplaceAllString(m[0][0], "\\${1}") +`([ \D]|$)`)
		s = rx.ReplaceAllString(s, "${1}" + calc(m[0][1]) + "${2}")

	}
	n := 0
	o := "+"
	for _, part := range strings.Split(s, " ") {
		switch part {
		case "+", "*":
			o = part

		default:
			num, _ := strconv.Atoi(part)
			if o == "+" {
				n += num
				continue
			}
			if o == "*" {
				n *= num
				continue

			}

		}
	}
	return fmt.Sprint(n)
}

func main() {
	input := strings.Split(strings.Trim(string(fileLoader.LoadFile("180-input")), "\n"), "\n")

	re := regexp.MustCompile(`\(([0-9+* ]*)\)`)
	r := 0
	for _, line := range input {
		for m := re.FindStringSubmatch(line); len(m) != 0; m = re.FindStringSubmatch(line) {
			line = strings.ReplaceAll(line, m[0], calc(m[1]))
		}
		res, _ := strconv.Atoi(calc(line))

		r += res
	}
	fmt.Println(r)

}
