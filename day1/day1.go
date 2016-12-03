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

	intersects := math.Abs(float64(c)) > 0.01

	if !intersects {
		return intersects, Coords{0, 0}
	}

	a := l1.a.x*l1.b.y - l1.a.y*l1.b.x
	b := l2.a.x*l2.b.y - l2.a.y*l2.b.x

	x := (a*x34 - b*x12) / c
	y := (a*y34 - b*y12) / c

	icoords := Coords{x, y}

	if outside_line(icoords, l1) || outside_line(icoords, l2) {
		return false, Coords{0, 0}
	}

	return intersects, icoords
}

func outside_line(coords Coords, line Line) bool {
	return outside(coords.x, line.a.x, line.b.x) ||
		outside(coords.y, line.a.y, line.b.y)
}

func outside(test, i1, i2 int) bool {
	low := min(i1, i2)
	high := max(i1, i2)
	return test < low || test > high
}

func min(i1, i2 int) int {
	return int(math.Min(float64(i1), float64(i2)))
}

func max(i1, i2 int) int {
	return int(math.Max(float64(i1), float64(i2)))
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
		coords.y += int(dist)
		break
	case E:
		coords.x += int(dist)
		break
	case S:
		coords.y -= int(dist)
		break
	case W:
		coords.x -= int(dist)
		break
	}
}

type Lines []Line

func (lines *Lines) add(line Line) {
	*lines = append(*lines, line)
}

func (lines *Lines) intersects(line Line) (bool, Coords) {
	for _, l := range *lines {
		i, coords := l.intersects(line)
		if i {
			return true, coords
		}
	}
	return false, Coords{0, 0}
}

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
	lines := make(Lines, 0)
	for val := range ch {
		prevcoord := coords
		facing = make_turn(facing, val.turn)
		make_move(&coords, facing, val.dist)
		line := Line{prevcoord, coords}
		ist, ic := lines.intersects(line)
		if ist {
			if prevcoord.x != ic.x || prevcoord.y != ic.y {
				coords = ic
				break
			}
		}
		lines.add(line)
	}
	fmt.Printf("Final coords: %+v\n", coords)
	fmt.Printf("Distance: %d\n", coords.x+coords.y)
}
