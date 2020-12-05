package main

import (
	"fmt"
	"github.com/poltergeist/aoc2020/fileLoader"
	"regexp"
	"strconv"
	"strings"
)

func byrvalidator(s string) bool {
	n, _ := strconv.Atoi(s)
	return n >= 1920 && n <= 2002
}

func iyrvalidator(s string) bool {
	n, _ := strconv.Atoi(s)
	return n >= 2010 && n <= 2020
}

func eyrvalidator(s string) bool {
	n, _ := strconv.Atoi(s)
	return n >= 2020 && n <= 2030
}

func hgtvalidator(s string) bool {
	r, _ := regexp.Compile(`(\d*)(cm|in)`)
	m := r.FindStringSubmatch(s)
	if len(m) != 3 {
		return false
	}
	switch m[2] {
	case "cm":
		s, _ := strconv.Atoi(m[1])
		return s >= 150 && s <= 193
	case "in":
		s, _ := strconv.Atoi(m[1])
		return s >= 59 && s <= 76

	}
	return false
}

func hclvalidator(s string) bool {
	r, _ := regexp.Compile(`^#[a-f,0-9]{6}$`)
	return r.MatchString(s)
}

func eclvalidator(s string) bool {
	return s == "amb" || s == "blu" || s == "brn" || s == "gry" || s == "grn" || s == "hzl" || s == "oth"
}

func pidvalidator(s string) bool {
	r, _ := regexp.Compile(`^[0-9]{9}$`)
	return r.MatchString(s)
}

func main() {
	required := map[string]func(s string) bool{
		"byr": byrvalidator,
		"iyr": iyrvalidator,
		"eyr": eyrvalidator,
		"hgt": hgtvalidator,
		"hcl": hclvalidator,
		"ecl": eclvalidator,
		"pid": pidvalidator,
	}
	input := strings.Split(string(fileLoader.LoadFile("./040-input")), "\n")
	valid := 0
	rf := []string{}
	for _, l := range input {
		if l == "" {
			if len(rf) >= len(required) {
				valid = valid + 1
			}
			rf = []string{}
			continue
		}
		for _, x := range strings.Split(l, " ") {
			f := strings.Split(x, ":")
			if fc, ok := required[f[0]]; ok {
				if fc(f[1])  {
					rf = append(rf, f[0])
				}
			}
		}
	}
	fmt.Println(valid)
}
