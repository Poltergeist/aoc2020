package day02

import (
	"strconv"
	"strings"
)

func checkEntry(i string) bool {
	if i == "" {
		return false
	}
	j := strings.Split(i, ":")
	k := strings.Split(j[0], " ")
	min, _ := strconv.Atoi(strings.Split(k[0], "-")[0])
	max, _ := strconv.Atoi(strings.Split(k[0], "-")[1])

	count := strings.Count(j[1], k[1])

	return count >= min && count <= max
}

func checkEntryPosition(i string) bool {
	if i == "" {
		return false
	}
	j := strings.Split(i, ":")
	k := strings.Split(j[0], " ")
	a, _ := strconv.Atoi(strings.Split(k[0], "-")[0])
	b, _ := strconv.Atoi(strings.Split(k[0], "-")[1])

	x := string(j[1][a]) == k[ 1 ]
	y := string(j[1][b]) == k[ 1 ]

	return (x || y) && !(x && y)
}

func Check(list []string) int {
	valid := 0
	for _, i := range list {
		if checkEntry(i) {
			valid = valid+ 1
		}
	}

	return valid
}

func CheckPosition(list []string) int {
	valid := 0
	for _, i := range list {
		if checkEntryPosition(i) {
			valid = valid+ 1
		}
	}

	return valid
}
