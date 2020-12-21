package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/poltergeist/aoc2020/fileLoader"
)

func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}
func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func main() {
	input := strings.Split(strings.Trim(string(fileLoader.LoadFile("210-input")), "\n"), "\n")
	al := map[string][]string{}
	i := []string{}
	rx := regexp.MustCompile(`\(contains ([a-zA-Z, ]*)\)`)
	for _, x := range input {
		xa := strings.Split(strings.Trim(rx.FindStringSubmatch(x)[1], " "), ", ")
		xi := x[:strings.Index(x, " (")]
		i = append(i, strings.Split(strings.Trim(xi, " "), " ")...)
		for _, a := range xa {
			if _, ok := al[a]; !ok {
				al[a] = []string{}
			}
			al[a] = append(al[a], xi)
		}

	}
	for a, l := range al {
		cs := strings.Split(strings.Trim(l[0], " "), " ")
		for _, le := range l[1:] {
			t := append([]string{}, cs...)
			for _, c := range cs {
				if strings.Index(le, c) == -1 && Index(t, c) != -1 {
					t = append(t[:Index(t, c)], t[Index(t, c)+1:]...)
				}
			}
			cs = append([]string{}, t...)
		}
		al[a] = cs
	}
	at := map[string]string{}
	oal := []string{}
	for len(al) > 0 {
		for a, l := range al {
			l = Filter(l, func(s string) bool {
				return Index(l, s) != -1 && Index(oal, s) == -1
			})
			if len(l) == 1 {
				at[a] = l[0]
				oal = append(oal, l...)
				delete(al, a)
			}
		}
	}
	i = Filter(i, func(s string) bool {
		return Index(oal, s) == -1
	})
	fmt.Println(at, "||", al)
	fmt.Println(len(i))
}
