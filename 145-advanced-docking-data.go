package main

import (
	"fmt"
	"github.com/poltergeist/aoc2020/fileLoader"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input := strings.Split(strings.Trim(string(fileLoader.LoadFile("140-input")), "\n"), "\n")
// 	input = strings.Split(strings.Trim(`mask = 000000000000000000000000000000X1001X
// mem[42] = 100
// mask = 00000000000000000000000000000000X0XX
// mem[26] = 1
// `, "\n"), "\n")

	re, _ := regexp.Compile(`(mask|mem)(?:\[)?(\d*)?(?:\])? = ([X0-9]*)`)
	mem := map[int64]int{}

	l := 0
	mask := ""
	for _, x := range input {
		m := re.FindStringSubmatch(x)
		m = m[1:]
		switch m[0] {
		case "mask":
			l = len(m[2])
			mask = strings.Trim(m[2], " ")

		case "mem":
			ml, _ := strconv.Atoi(m[1])
			v, _ := strconv.Atoi(m[2])
			bmls := strings.Repeat("0", 100) + strconv.FormatInt(int64(ml), 2)
			bmls = bmls[len(bmls)-l:]
			mz := []string{""}
			for _, mc := range mask {
				switch string(mc) {
				case "X":
					mz = append(mz, mz...)
					for k, _ := range mz {
						if len(mz)/2 > k {
							mz[k] += "0"
							continue
						}
						mz[k] += "1"
					}
				case "0":
					for k, _ := range mz {
						mz[k] += string(bmls[len(mz[k])])

					}
				case "1":
					for k, _ := range mz {
						mz[k] += "1"
					}
				}
			}
			for _, mv := range mz {
				d, _ := strconv.ParseInt(mv, 2, 64)
				mem[d] = v
			}
		}
	}

	sum := 0
	for _, a := range mem {
		sum += a
	}
	fmt.Println(sum)
}
