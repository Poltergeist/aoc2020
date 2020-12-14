package main

import (
	"fmt"
	"regexp"
	"strings"
	"strconv"
	"github.com/poltergeist/aoc2020/fileLoader"
)

func main() {
	input := strings.Split(strings.Trim(string(fileLoader.LoadFile("140-input")), "\n"), "\n")
	re, _ := regexp.Compile(`(mask|mem)(?:\[)?(\d*)?(?:\])? = ([X0-9]*)`)
	mem := map[int]int64{}

	l := 0
	mask := ""
	for _, x := range input {
		m := re.FindStringSubmatch(x)
		m = m[1:]
		switch m[0] {
		case "mask":
			l = len(m[2])
			mask = m[2]
		case "mem":
			ml, _ := strconv.Atoi(m[1])
			x, _ := strconv.Atoi(m[2])
			bs := strings.Repeat("0", 100) + strconv.FormatInt(int64(x), 2)
			bs = bs[len(bs)-l:]
			y := ""
			for k, b := range bs {
				if string(mask[k]) != "X" {
					y += string(mask[k])
					continue
				}
				y += string(b)
			}
			mem[ml], _ = strconv.ParseInt(y, 2, 64)

		}
	}

	sum := int64(0)
	for _, a := range mem {
		sum += a
	}
	fmt.Println(sum)
}
