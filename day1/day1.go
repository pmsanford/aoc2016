package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	L = -1
	R = 1
)

const (
	N int = 0
	E int = 1
	S int = 2
	W int = 3
)

type Direction struct {
	turn int
	dist uint
}

type Coords struct {
	x int
	y int
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

func parse_input(filename string, ch chan Direction) {
	defer close(ch)
	dat, err := ioutil.ReadFile(filename)
	check(err)
	istr := string(dat)
	dirs := strings.Split(istr, ", ")
	for i := range dirs {
		dirstr := strings.Trim(dirs[i], "\r\n ")
		if len(dirstr) < 2 {
			continue
		}
		turnid := dirstr[0]
		turn := R
		if turnid == 76 {
			turn = L
		}
		count, err := strconv.ParseInt(dirstr[1:], 10, 64)
		check(err)
		direction := Direction{turn, uint(count)}
		ch <- direction
	}
}

func make_turn(facing, turn_dir int) int {
	facing = facing + turn_dir
	if facing == -1 {
		facing = 3
	}
	if facing == 4 {
		facing = 0
	}
	return facing
}

func make_move(coords *Coords, facing int, dist uint) {
	switch facing {
	case N:
		coords.x += int(dist)
		break
	case E:
		coords.y += int(dist)
		break
	case S:
		coords.x -= int(dist)
		break
	case W:
		coords.y -= int(dist)
		break
	}
}

func main() {
	ch := make(chan Direction)
	go parse_input("input.txt", ch)
	facing := N
	coords := Coords{0, 0}
	for val := range ch {
		facing = make_turn(facing, val.turn)
		make_move(&coords, facing, val.dist)
	}
	fmt.Println(coords)
}
