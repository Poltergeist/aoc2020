package misc

import (
	"strconv"
)

func ConvertStringsToInt(list []string) []int {
	var returnValue = []int{}

	for _, i := range list {
		j, err := strconv.Atoi(i)
		if err != nil {
			continue
		}
		returnValue = append(returnValue, j)
	}
	return returnValue
}
