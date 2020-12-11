package main

import (
	"fmt"
	"github.com/poltergeist/aoc2020/fileLoader"
	"reflect"
	"strings"
)

func main() {
	temp := strings.Split(strings.Trim(string(fileLoader.LoadFile("110-input")), "\n"), "\n")
// 	temp = strings.Split(strings.Trim(`L.LL.LL.LL
// LLLLLLL.LL
// L.L.L..L..
// LLLL.LL.LL
// L.LL.LL.LL
// L.LLLLL.LL
// ..L.L.....
// LLLLLLLLLL
// L.LLLLLL.L
// L.LLLLL.LL
// `, "\n"), "\n")
	input := [][]string{}
	for _, v := range temp {
		input = append(input, strings.Split(v, ""))
	}

	for {
		next := [][]string{}
		for _, s := range input {
			next = append(next, strings.Split(strings.Join(s, ""), ""))
		}
		for row, x := range input {
			for col, y := range x {
				o := 0
				for i := -1; i <= 1; i++ {
					if row+i >= 0 && row+i < len(input) {

						for j := -1; j <= 1; j++ {
							if col+j >= 0 && col+j < len(x) {
								if !(i == 0 && j == 0) {
									for cx, cy := i, j; col+cy >= 0 && col+cy < len(x) && row+cx >= 0 && row+cx < len(input); cx, cy = cx+i, cy+j {
										if input[row+cx][col+cy] != "." {
											if input[row+cx][col+cy] == "#" {
												o++
											}
											break
										}
									}
								}
							}

						}

					}

				}
				switch y {
				case "L":
					if o == 0 {
						next[row][col] = "#"
					}

				case "#":
					if o >= 5 {
						next[row][col] = "L"
					}
				}

			}
		}
		if reflect.DeepEqual(input, next) {
			c := 0
			for _, a := range next {
				for _, b := range a {
					if b == "#" {
						c++
					}
				}
			}
			fmt.Println(c)
			return
		}
		input = [][]string{}
		for _, s := range next {
			input = append(input, strings.Split(strings.Join(s, ""), ""))
		}
	}
}
