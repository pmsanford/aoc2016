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
		fmt.Printf("%X", key)
	}
}

func lookup_move(key, dir int) int {
	m := Lookup{
		1: KeyLookup{
			U: 1,
			R: 1,
			L: 1,
			D: 3,
		},
		2: KeyLookup{
			U: 2,
			R: 3,
			L: 2,
			D: 6,
		},
		3: KeyLookup{
			U: 1,
			R: 4,
			L: 2,
			D: 7,
		},
		4: KeyLookup{
			U: 4,
			R: 4,
			L: 3,
			D: 8,
		},
		5: KeyLookup{
			U: 5,
			R: 6,
			L: 5,
			D: 5,
		},
		6: KeyLookup{
			U: 2,
			R: 7,
			L: 5,
			D: 10,
		},
		7: KeyLookup{
			U: 3,
			R: 8,
			L: 6,
			D: 11,
		},
		8: KeyLookup{
			U: 4,
			R: 9,
			L: 7,
			D: 12,
		},
		9: KeyLookup{
			U: 9,
			R: 9,
			L: 8,
			D: 9,
		},
		10: KeyLookup{
			U: 6,
			R: 11,
			L: 10,
			D: 10,
		},
		11: KeyLookup{
			U: 7,
			R: 12,
			L: 10,
			D: 13,
		},
		12: KeyLookup{
			U: 8,
			R: 12,
			L: 11,
			D: 12,
		},
		13: KeyLookup{
			U: 11,
			R: 13,
			L: 13,
			D: 13,
		},
	}
	return m[key][dir]
}
