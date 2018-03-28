package main

import (
	"math"
	"fmt"
)

type Point struct {
	X, Y float64
}

func (p Point)Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y - p.Y)
}

type Path []Point

func (p Path)Distance() float64 {
	return 0
}

func (p Path)LenPoints() int {
	return len(p)
}

type MyString string

func (s MyString)substr(i, j int) string {
	var s1 = string(s)
	return s1[i:j]
}

func main() {
	p := Point{1, 2}
	q := Point{4, 6}
	fmt.Println(p.Distance(q))

	s := MyString("hello")
	fmt.Println(s.substr(1, 2))

	var a struct{
		Id int
		X float64
		Y float64
	}
	a.Id = 1
}
