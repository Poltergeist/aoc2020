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

func Check(list []string) int {
	valid := 0
	for _, i := range list {
		if checkEntry(i) {
			valid = valid+ 1
		}
	}

	return valid
}
