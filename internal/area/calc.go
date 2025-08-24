package area

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float64
}

func NewCircle(radius float64) Circle {
	return Circle{
		Radius: radius,
	}
}

func (c Circle) Area() float64 {
	return (math.Pi * math.Pow(float64(c.Radius), 2))
}

type Rectangle struct {
	Height float64
	Width  float64
}

func NewRectangle(height, width float64) Rectangle {
	return Rectangle{
		Height: height,
		Width:  width,
	}
}

func (r Rectangle) Area() float64 {
	return (r.Height * r.Width)
}

type Geometric interface {
	Area() float64
}

func Calculator(shapes ...Geometric) {
	totalArea := float64(0)
	for i, v := range shapes {
		area := v.Area()
		fmt.Printf("%d : %.2f\t", i+1, area)
		totalArea = totalArea + area
	}

	fmt.Printf("\nTotal area : %.2f\n", totalArea)
}
