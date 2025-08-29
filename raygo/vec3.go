package main

import "math"

type Vec3 struct {
	X float64
	Y float64
	Z float64
}

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
