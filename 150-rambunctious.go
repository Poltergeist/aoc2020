package main

import (
	"fmt"
)

func main() {
	input := []int{0,12,6,13,20,1,17}
	s := map[int][]int{}

	for k, v := range input {
		s[v] = []int{k}
	}

	l := input[len(input)-1]
	for c := len(input); c < 2020; c++ {
		if len(s[l]) == 1 {
			s[0] = append([]int{c}, s[0]...)
			l = 0
			continue
		}
		n := s[l][0] - s[l][1]
		if len(s[n]) == 0 {
			s[n] = []int{c}
			l = n
			continue
		}
		s[n] = []int{c, s[n][0]}
		l = n

	}
	fmt.Println(l)

}
