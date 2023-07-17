package main

import (
	"fmt"
)

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	w := Wheel{
		Spokes: 4,
		Circle: Circle{
			Point:  Point{X: 3, Y: 4},
			Radius: 4,
		},
	}

	fmt.Printf("%v\n", w)
	fmt.Printf("%#v\n", w)
}
