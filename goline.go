package goline

import (
	"fmt"
	"math"
)

// http://en.wikipedia.org/wiki/Linear_equation

const epsilon = 1e-6

func Equal(a, b float64) bool {
	return (math.Abs(a-b) < epsilon)
}

type Point struct {
	X, Y float64
}

// ax + by + c = 0
type Line struct {
	a, b, c float64
}

func NewLine(a, b, c float64) *Line {
	return &Line{a, b, c}
}

func NewLineFromPoints(m1, m2 Point) *Line {

	dX := m2.X - m1.X
	dY := m2.Y - m1.Y

	return &Line{
		a: dY,
		b: -dX,
		c: dX*m1.Y - dY*m1.X,
	}
}

func (this *Line) XFromY(y float64) (x float64) {

	x = -(this.b*y + this.c) / this.a
	return
}

func (this *Line) YFromX(x float64) (y float64) {

	y = -(this.a*x + this.c) / this.b
	return
}

func (this *Line) String() string {
	return fmt.Sprintf("(%g)*x + (%g)*y + (%g) = 0", this.a, this.b, this.c)
}
