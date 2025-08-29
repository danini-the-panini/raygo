package main

import (
	"math"
	"math/rand"
)

type Vec3 struct {
	X float64
	Y float64
	Z float64
}

var ZERO3 = Vec3{0.0, 0.0, 0.0}

func Origin3() Vec3 {
	return Vec3{X: 0.0, Y: 0.0, Z: 0.0}
}

func (v *Vec3) negate() *Vec3 {
	v.X = -v.X
	v.Y = -v.Y
	v.Z = -v.Z
	return v
}

func (v Vec3) inverse() Vec3 {
	return Vec3{X: -v.X, Y: -v.Y, Z: -v.Z}
}

func (v Vec3) get(i int) float64 {
	switch i {
	case 0:
		return v.X
	case 1:
		return v.Y
	case 2:
		return v.Z
	default:
		return 0.0
	}
}

func (v *Vec3) set(i int, a float64) *Vec3 {
	switch i {
	case 0:
		v.X = a
	case 1:
		v.Y = a
	case 2:
		v.Z = a
	}
	return v
}

func (v *Vec3) add(u Vec3) *Vec3 {
	v.X += u.X
	v.Y += u.Y
	v.Z += u.Z
	return v
}

func (v Vec3) plus(u Vec3) Vec3 {
	return Vec3{X: v.X + u.X, Y: v.Y + u.Y, Z: v.Z + u.Z}
}

func (v *Vec3) sub(u Vec3) *Vec3 {
	v.X -= u.X
	v.Y -= u.Y
	v.Z -= u.Z
	return v
}

func (v Vec3) minus(u Vec3) Vec3 {
	return Vec3{X: v.X - u.X, Y: v.Y - u.Y, Z: v.Z - u.Z}
}

func (v *Vec3) scale(t float64) *Vec3 {
	v.X *= t
	v.Y *= t
	v.Z *= t
	return v
}

func (v Vec3) mul(u Vec3) Vec3 {
	return Vec3{v.X * u.X, v.Y * u.Y, v.Z * u.Z}
}

func (v Vec3) times(t float64) Vec3 {
	return Vec3{X: v.X * t, Y: v.Y * t, Z: v.Z * t}
}

func (v *Vec3) div(t float64) *Vec3 {
	v.scale(1.0 / t)
	return v
}

func (v Vec3) divBy(t float64) Vec3 {
	return v.times(1.0 / t)
}

func (v Vec3) len() float64 {
	return math.Sqrt(v.lenSq())
}

func (v Vec3) lenSq() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

func (v Vec3) dot(u Vec3) float64 {
	return v.X*u.X + v.Y*u.Y + v.Z*u.Z
}

func (v Vec3) cross(u Vec3) Vec3 {
	return Vec3{
		X: v.Y*u.Z - v.Z*u.Y,
		Y: v.Z*u.X - v.X*u.Z,
		Z: v.X*u.Y - v.Y*u.X,
	}
}

func (v Vec3) unit() Vec3 {
	return v.divBy(v.len())
}

func (v *Vec3) normalize() *Vec3 {
	return v.div(v.len())
}

func (v Vec3) nearZero() bool {
	var s = 1e-8
	return math.Abs(v.X) < s && math.Abs(v.Y) < s && math.Abs(v.Z) < s
}

func (v Vec3) reflect(n Vec3) Vec3 {
	return v.minus(n.times(2.0 * v.dot(n)))
}

func SampleSquare() Vec3 {
	return Vec3{rand.Float64() - 0.5, rand.Float64() - 0.5, 0.0}
}

func RandVec3() Vec3 {
	return Vec3{rand.Float64(), rand.Float64(), rand.Float64()}
}

func RandUnit() Vec3 {
	var i = Interval{-1.0, 1.0}
	for {
		var p = i.randVec3()
		var lensq = p.lenSq()
		if 1e-160 < lensq && lensq <= 1.0 {
			return *p.div(math.Sqrt(lensq))
		}
	}
}

func RandHemi(normal Vec3) Vec3 {
	var on_unit_sphere = RandUnit()
	if on_unit_sphere.dot(normal) < 0.0 { // In the same hemisphere as the normal
		on_unit_sphere.negate()
	}
	return on_unit_sphere
}
