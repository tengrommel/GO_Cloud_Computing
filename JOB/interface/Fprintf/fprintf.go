package main

import (
	"fmt"
	"net/http"
	"math"
)

func handle(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "hello\n")
}

// 名字： IInstance
// 方法： Instance() float64
type IInstance interface {
	Instance() float64
}

type Point struct {
	X float64
	Y float64
}

type Path []*Point

func (p Path)Instance() float64 {
	return 0.1
}

func (p Point)Distance2Point(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y- q.X)
}

func (p Point)Distance() float64 {
	return math.Hypot(p.X, p.Y)
}


func (p Path)Distance() float64{
	var sum float64
	for i:=0;i<len(p)-1;i++{
		sum += p[i].Distance2Point(*p[i+1])
	}
	return sum
}

type IDistance interface {
		Distance() float64
}

func main() {
	var path Path
	path = make([]*Point, 3)
	p1 := &Point{X:1, Y:2}
	p2 := &Point{X:3, Y:4}
	p3 := &Point{X:5, Y:6}
	path[0] = p1
	path[1] = p2
	path[2] = p3

	var i IDistance
	i = p1
	fmt.Println(i.Distance())
	i = p2
	fmt.Println(i.Distance())
	i = path
	fmt.Println(i.Distance())
	fmt.Println(path)
	fmt.Println(p1)
}
