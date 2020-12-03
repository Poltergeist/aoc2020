package day03

import (
)

type Move struct {
	R int
	D int
}

func TreeCount(treeMap []string, m Move) int {
	r := m.R
	width := len(treeMap[0])
	count := 0
	for k, x := range treeMap {
		if k % m.D != 0 {
			continue
		}
		key := ((k/m.D * r) % width)

		if string(x[key]) == "#" {
			count = count + 1
		}
	}
	return count
}
