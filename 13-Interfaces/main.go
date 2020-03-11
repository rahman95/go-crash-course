package main

import (
	"fmt"
	"math"
)

// Shape interface
type Shape interface {
	area() float64
}

// Circle Struct
type Circle struct {
	radius float64
}

// Rectangle Struct
type Rectangle struct {
	height, width float64
}

func (circle Circle) area() float64 {
	return math.Pi * circle.radius * circle.radius
}

func (rectangle Rectangle) area() float64 {
	return rectangle.width * rectangle.height
}

func getArea(shape Shape) float64 {
	return shape.area()
}

func main() {
	circle := Circle{radius: 10}
	rectangle := Rectangle{height: 5, width: 10}

	fmt.Println(circle, rectangle)
	fmt.Println(getArea(circle), getArea(rectangle))
}
