package fileLoader

import (
	"io/ioutil"
)


func check(e error) {
	if e != nil {
		panic(e)
	}
}

func LoadFile(path string) []byte {
	dat, err := ioutil.ReadFile(path)
	check(err)
	return dat
}
