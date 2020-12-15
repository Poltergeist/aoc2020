package main

import (
	"fmt"
)

func main() {
	input := []int{0, 12, 6, 13, 20, 1, 17}
	l := input[len(input)-1]
	n := 30000000
	ls := make([]int, n)
	for k, v := range input {
		ls[v] = k + 1
	}

	for c := len(input); c < n; c++ {
		t := 0
		if ls[l] != 0 {
			t = c - ls[l]
		}
		ls[l] = c
		l = t

	}
	fmt.Println(l)

}
