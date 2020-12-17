package main

import (
	"fmt"
	"strings"
	"github.com/poltergeist/aoc2020/fileLoader"
)

func setState(state map[int]map[int]map[int]bool, x int, y int, z int, value bool) map[int]map[int]map[int]bool {
	for _, xd := range []int{-1, 0, 1} {
		for _, yd := range []int{-1, 0, 1} {
			for _, zd := range []int{-1, 0, 1} {
				if _, ok := state[x+xd]; !ok {
					state[x+xd] = make(map[int]map[int]bool)
				}
				if _, ok := state[x+xd][y+yd]; !ok {
					state[x+xd][y+yd] = make(map[int]bool)
				}
				if _, ok := state[x+xd][y+yd][z+zd]; !ok {
					state[x+xd][y+yd][z+zd] = false

				}
			}
		}
	}
	state[x][y][z] = value
	return state
}

func main() {
	input := strings.Split(strings.Trim(string(fileLoader.LoadFile("170-input")), "\n"),"\n")
// 	input = strings.Split(strings.Trim(`.#.
// ..#
// ###`, "\n"), "\n")

	state := make(map[int]map[int]map[int]bool)

	for x, row := range input {
		for y, value := range strings.Split(row, "") {
			state = setState(state, x, y, 0, value == "#")
		}
	}

	for c := 6; c > 0; c-- {
		next := make(map[int]map[int]map[int]bool)
		const MaxUint = ^uint(0)
		const MaxInt = int(MaxUint >> 1)
		const MinInt = -MaxInt - 1

		xMin := MaxInt
		xMax := MinInt
		yMin := MaxInt
		yMax := MinInt
		zMin := MaxInt
		zMax := MinInt

		for x, _ := range state {
			if x < xMin {
				xMin = x
			}
			if x > xMax {
				xMax = x
			}
			for y, _ := range state[x] {
				if y < yMin {
					yMin = y
				}
				if y > yMax {
					yMax = y
				}
				for z, _ := range state[x][y] {
					if z < zMin {
						zMin = z
					}
					if z > zMax {
						zMax = z
					}
				}
			}
		}

		for x := xMin; x <= xMax; x++ {
			for y := yMin; y <= yMax; y++ {
				for z := zMin; z <= zMax; z++ {
					value := state[x][y][z]
					active := 0
					// Check
					for _, xd := range []int{-1, 0, 1} {
						for _, yd := range []int{-1, 0, 1} {
							for _, zd := range []int{-1, 0, 1} {
								if xd == 0 && yd == 0 && zd == 0 {
									continue
								}
								if v, ok := state[x+xd][y+yd][z+zd]; ok && v {
									active++
								}

							}
						}
					}
					next = setState(next, x, y, z, (value && (active == 2 || active == 3)) || (!value && active == 3))

				}

			}
		}
		state = next
	}

	a := 0
	for _, row := range state {
		for _, col := range row {
			for _, value := range col {
				if value == true {
					a++
				}
			}
		}
	}

	fmt.Println(a)

}
