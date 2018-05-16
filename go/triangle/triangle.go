// Package triangle should have a package comment that summarizes what it's about.
package triangle

import "math"

type Kind string

const (
	NaT Kind = "NaT" // not a triangle
	Equ      = "Equ" // equilateral
	Iso      = "Iso" // isosceles
	Sca      = "Sca" // scalene
)

func KindFromSides(a, b, c float64) Kind {
	for _, side := range []float64{a, b, c} {
		if side <= 0 || math.IsNaN(side) || math.IsInf(side, 1) || math.IsInf(side, -1) {
			return NaT
		}
	}
	switch {
	case a+b < c || a+c < b || b+c < a:
		return NaT
	case a == b && b == c:
		return Equ
	case a == b || a == c || b == c:
		return Iso
	default:
		return Sca
	}
}