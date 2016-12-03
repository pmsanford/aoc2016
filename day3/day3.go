package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Triangle struct {
	a int
	b int
	c int
}

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

func parse_line(line string) Triangle {
	a := 0
	b := 0
	c := 0
	fmt.Sscanf(line, "%d %d %d", &a, &b, &c)
	return Triangle{a, b, c}
}

func is_valid_triangle(triangle Triangle) bool {
	return triangle.a+triangle.b > triangle.c &&
		triangle.a+triangle.c > triangle.b &&
		triangle.b+triangle.c > triangle.a
}

func transpose_triangles(t1, t2, t3 Triangle) (Triangle, Triangle, Triangle) {
	nt1 := Triangle{t1.a, t2.a, t3.a}
	nt2 := Triangle{t1.b, t2.b, t3.b}
	nt3 := Triangle{t1.c, t2.c, t3.c}

	return nt1, nt2, nt3
}

func main() {
	lines := parse_input("input.txt")
	valid := 0
	for i := 0; i < len(lines); i += 3 {
		tri1 := parse_line(lines[i])
		tri2 := parse_line(lines[i+1])
		tri3 := parse_line(lines[i+2])

		tri1, tri2, tri3 = transpose_triangles(tri1, tri2, tri3)
		if is_valid_triangle(tri1) {
			valid += 1
		}
		if is_valid_triangle(tri2) {
			valid += 1
		}
		if is_valid_triangle(tri3) {
			valid += 1
		}
	}
	fmt.Printf("Valid: %d", valid)
}
