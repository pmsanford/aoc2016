package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	L uint = iota
	R
)

type Direction struct {
	turn uint
	dist uint
}

func (dir Direction) String() string {
	turn := "L"
	if dir.turn == R {
		turn = "R"
	}
	return fmt.Sprintf("Turn: %s Distance: %d", turn, dir.dist)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parse_input(filename string) []Direction {
	dat, err := ioutil.ReadFile(filename)
	check(err)
	istr := string(dat)
	dirs := strings.Split(istr, ", ")
	dirlist := make([]Direction, 100)
	for i := range dirs {
		turnid := dirs[i][0]
		turn := R
		if turnid == 76 {
			turn = L
		}
		count, _ := strconv.ParseInt(dirs[i][1:], 10, 64)
		direction := Direction{turn, uint(count)}
		dirlist = append(dirlist, direction)
	}
	return dirlist
}

func main() {
	dirlist := parse_input("input.txt")
	for _, val := range dirlist {
		fmt.Println(val)
	}
}
