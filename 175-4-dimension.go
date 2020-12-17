package main

import (
	"fmt"
	"github.com/poltergeist/aoc2020/fileLoader"
	"strings"
)

func setState(state map[int]map[int]map[int]map[int]bool, x int, y int, z int, w int, value bool) map[int]map[int]map[int]map[int]bool {
	for _, xd := range []int{-1, 0, 1} {
		for _, yd := range []int{-1, 0, 1} {
			for _, zd := range []int{-1, 0, 1} {
				for _, wd := range []int{-1, 0, 1} {
					if _, ok := state[x+xd]; !ok {
						state[x+xd] = make(map[int]map[int]map[int]bool)
					}
					if _, ok := state[x+xd][y+yd]; !ok {
						state[x+xd][y+yd] = make(map[int]map[int]bool)
					}
					if _, ok := state[x+xd][y+yd][z+zd]; !ok {
						state[x+xd][y+yd][z+zd] = make(map[int]bool)

					}
					if _, ok := state[x+xd][y+yd][z+zd][w+wd]; !ok {
						state[x+xd][y+yd][z+zd][w+wd] = false

					}
				}
			}
		}
	}
	state[x][y][z][w] = value
	return state
}

func main() {
	input := strings.Split(strings.Trim(string(fileLoader.LoadFile("170-input")), "\n"), "\n")
	// 	input = strings.Split(strings.Trim(`.#.
	// ..#
	// ###`, "\n"), "\n")

	state := make(map[int]map[int]map[int]map[int]bool)

	for x, row := range input {
		for y, value := range strings.Split(row, "") {
			state = setState(state, x, y, 0, 0, value == "#")
		}
	}

	for c := 6; c > 0; c-- {
		next := make(map[int]map[int]map[int]map[int]bool)

		for x, _ := range state {
			for y, _ := range state[x] {
				for z, _ := range state[x][y] {
					for w, _ := range state[x][y][z] {
						value := state[x][y][z][w]
						active := 0
						// Check
						for _, xd := range []int{-1, 0, 1} {
							for _, yd := range []int{-1, 0, 1} {
								for _, zd := range []int{-1, 0, 1} {
									for _, wd := range []int{-1, 0, 1} {
										if xd == 0 && yd == 0 && zd == 0 && wd == 0 {
											continue
										}
										if v, ok := state[x+xd][y+yd][z+zd][w+wd]; ok && v {
											active++
										}
									}

								}
							}
						}
						next = setState(next, x, y, z, w, (value && (active == 2 || active == 3)) || (!value && active == 3))

					}
				}

			}
		}
		state = next
	}

	a := 0
	for _, row := range state {
		for _, col := range row {
			for _, z := range col {
				for _, value := range z {
					if value == true {
						a++
					}
				}
			}
		}
	}

	fmt.Println(a)

}
