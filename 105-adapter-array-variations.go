package main

import (
	"fmt"
	"github.com/poltergeist/aoc2020/fileLoader"
	"sort"
	"strconv"
	"strings"
)

func main() {
	temp := strings.Split(strings.Trim(string(fileLoader.LoadFile("100-input")), "\n"), "\n")
	input := []int{}
	for _, i := range temp {
		x, _ := strconv.Atoi(i)
		input = append(input, x)
	}
	sort.Ints(input)
	input = append([]int{0}, append(input, input[len(input)-1]+3)...)

	diff, memo := map[int]int{}, map[int]int{0: 1}
	for i, v := range input[1:] {
		diff[v-input[i]]++
		memo[v] = memo[v-1] + memo[v-2] + memo[v-3]
	}
	fmt.Println(memo[input[len(input)-1]])
}
