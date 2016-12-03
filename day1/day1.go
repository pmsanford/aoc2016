package main

import (
	"fmt"
	"io/ioutil"
	"math"
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

type Line struct {
	a Coords
	b Coords
}

func (l1 Line) intersects(l2 Line) (bool, Coords) {
	x12 := l1.a.x - l1.b.x
	x34 := l2.a.x - l2.b.x
	y12 := l1.a.y - l1.b.y
	y34 := l2.a.y - l2.b.y

	c := x12*y34 - y12*x34

	fmt.Printf("C: %d\n", c)

	intersects := math.Abs(float64(c)) > 0.01

	if !intersects {
		return intersects, Coords{0, 0}
	}

	a := l1.a.x*l1.b.y - l1.a.y*l1.b.x
	b := l2.a.x*l2.b.y - l2.a.y*l2.b.x

	x := (a*x34 - b*x12) / c
	y := (a*y34 - b*y12) / c

	return intersects, Coords{x, y}
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

type LineSet map[Line]bool

type DirSet map[Coords]bool

func (set DirSet) contains(dir Coords) bool {
	_, in := set[dir]
	return in
}

func (set *DirSet) add(dir Coords) {
	(*set)[dir] = true
}

func main() {
	ch := make(chan Direction)
	go parse_input("input.txt", ch)
	facing := N
	coords := Coords{0, 0}
	visited := make(DirSet)
	for val := range ch {
		facing = make_turn(facing, val.turn)
		make_move(&coords, facing, val.dist)
		if visited.contains(coords) {
			fmt.Printf("Found double-visit at %+v\n", coords)
			break
		} else {
			visited.add(coords)
		}
	}
	fmt.Printf("Final coords: %+v\n", coords)
	fmt.Printf("Distance: %d\n", coords.x+coords.y)

	a := Line{Coords{0, 0}, Coords{4, 4}}
	b := Line{Coords{2, 0}, Coords{2, 4}}
	c := Line{Coords{0, 0}, Coords{0, 2}}

	abi, abc := a.intersects(b)
	bci, _ := b.intersects(c)

	fmt.Printf("AB: %v at %+v, BC: %v", abi, abc, bci)
}
