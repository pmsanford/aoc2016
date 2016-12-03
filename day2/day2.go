package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	U int = 0
	R int = 1
	L int = 2
	D int = 3
)

type Lookup map[int]map[int]int
type KeyLookup map[int]int

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parse_input(filename string) []string {
	dat, err := ioutil.ReadFile(filename)
	check(err)
	istr := string(dat)
	lines := strings.Split(istr, "\n")
	return lines
}

func get_dir(letter int) int {
	switch letter {
	case 85:
		return U
	case 82:
		return R
	case 76:
		return L
	case 68:
		return D
	}
	return 0
}

func main() {
	lines := parse_input("input.txt")
	key := 5
	fmt.Printf("Key: ")
	for _, line := range lines {
		for _, val := range line {
			key = lookup_move(key, get_dir(int(val)))
		}
		fmt.Printf("%d", key)
	}
}

func lookup_move(key, dir int) int {
	m := Lookup{
		1: KeyLookup{
			U: 1,
			R: 2,
			L: 1,
			D: 4,
		},
		2: KeyLookup{
			U: 2,
			R: 3,
			L: 1,
			D: 5,
		},
		3: KeyLookup{
			U: 3,
			R: 3,
			L: 2,
			D: 6,
		},
		4: KeyLookup{
			U: 1,
			R: 5,
			L: 4,
			D: 7,
		},
		5: KeyLookup{
			U: 2,
			R: 6,
			L: 4,
			D: 8,
		},
		6: KeyLookup{
			U: 3,
			R: 6,
			L: 5,
			D: 9,
		},
		7: KeyLookup{
			U: 4,
			R: 8,
			L: 7,
			D: 7,
		},
		8: KeyLookup{
			U: 5,
			R: 9,
			L: 7,
			D: 8,
		},
		9: KeyLookup{
			U: 6,
			R: 9,
			L: 8,
			D: 9,
		},
	}
	return m[key][dir]
}
