package main

import (
	"fmt"
	"math"
)

type God struct {
	Name string
}

func (this God) WhoAmI() {
	fmt.Println("I am ", this.Name)
}

/* sprint-1 */
type Circle struct /* implements AreaCalculator */ {
	God
	Radius float32
}

func (this Circle) CalculateArea() float32 {
	return math.Pi * this.Radius * this.Radius
}

func (this Circle) CalculatePerimeter() float32 {
	return 2 * math.Pi * this.Radius
}

/* sprint-2 */
type Rectangle struct /* implements AreaCalculator */ {
	God
	Height float32
	Width  float32
}

func (this Rectangle) CalculateArea() float32 {
	return this.Height * this.Width
}

func (this Rectangle) CalculatePerimeter() float32 {
	return 2 * (this.Height + this.Width)
}

/* sprint-3 */
/* Utility Functions */
type AreaCalculator interface {
	CalculateArea() float32
}

func PrintArea(x AreaCalculator) {
	fmt.Println("Area :", x.CalculateArea())
}

type PerimeterCalculator interface {
	CalculatePerimeter() float32
}

func PrintPerimeter(x PerimeterCalculator) {
	fmt.Println("Perimeter :", x.CalculatePerimeter())
}

/* interface composition */
type ShapeCalculator interface {
	AreaCalculator
	PerimeterCalculator
}

func PrintShape(x ShapeCalculator) {
	PrintArea(x)
	PrintPerimeter(x)
}

func main() {
	c := Circle{Radius: 12, God: God{Name: "Circle"}}
	// fmt.Println("Area :", c.CalculateArea())
	/*
		PrintArea(c)
		PrintPerimeter(c)
	*/
	PrintShape(c)
	c.WhoAmI()

	r := Rectangle{Height: 10, Width: 12, God: God{Name: "Rectangle"}}
	// fmt.Println("Area :", r.CalculateArea())
	/*
		PrintArea(r)
		PrintPerimeter(r)
	*/
	PrintShape(r)
	r.WhoAmI()
}
