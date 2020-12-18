package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"github.com/poltergeist/aoc2020/fileLoader"
)

func calc(s string) string {
	parts := strings.Split(s, " ")
	n := 0
	o := "+"
	for _, part := range parts {
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
// 	input = strings.Split(strings.Trim(`1 + 2 * 3 + 4 * 5 + 6
// ((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2`, "\n"), "\n")

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
