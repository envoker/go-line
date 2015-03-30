package goline

import (
	"testing"
)

func TestLine(t *testing.T) {

	m1 := Point{10.9, -4.3}
	m2 := Point{-7.45, 13.1}

	line := NewLineFromPoints(m1, m2)
	t.Log(line)

	x1 := 10.0
	y := line.YFromX(x1)
	x2 := line.XFromY(y)

	if !Equal(x1, x2) {
		t.Error("x1 != x2")
	}
}
