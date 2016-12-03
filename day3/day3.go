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

func main() {
	lines := parse_input("input.txt")
	valid := 0
	for _, val := range lines {
		tri := parse_line(val)
		if is_valid_triangle(tri) {
			valid += 1
		}
	}
	fmt.Printf("Valid: %d", valid)
}
