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

func Negate(v *Vec3) *Vec3 {
	v.X = -v.X
	v.Y = -v.Y
	v.Z = -v.Z
	return v
}

func Neg(v Vec3) Vec3 {
	return Vec3{X: -v.X, Y: -v.Y, Z: -v.Z}
}

func Ref(v Vec3, i int) float64 {
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

func Set(v *Vec3, i int, a float64) *Vec3 {
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

func Add(v *Vec3, u Vec3) *Vec3 {
	v.X += u.X
	v.Y += u.Y
	v.Z += u.Z
	return v
}

func Plus(v Vec3, u Vec3) Vec3 {
	return Vec3{X: v.X + u.X, Y: v.Y + u.Y, Z: v.Z + u.Z}
}

func Sub(v *Vec3, u Vec3) *Vec3 {
	v.X -= u.X
	v.Y -= u.Y
	v.Z -= u.Z
	return v
}

func Minus(v Vec3, u Vec3) Vec3 {
	return Vec3{X: v.X - u.X, Y: v.Y - u.Y, Z: v.Z - u.Z}
}

func Scale(v *Vec3, t float64) *Vec3 {
	v.X *= t
	v.Y *= t
	v.Z *= t
	return v
}

func Times(v Vec3, t float64) Vec3 {
	return Vec3{X: v.X * t, Y: v.Y * t, Z: v.Z * t}
}

func Div(v *Vec3, t float64) *Vec3 {
	Scale(v, 1.0/t)
	return v
}

func DivBy(v Vec3, t float64) Vec3 {
	return Times(v, 1.0/t)
}

func Length(v Vec3) float64 {
	return math.Sqrt(LengthSq(v))
}

func LengthSq(v Vec3) float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

func Dot(v Vec3, u Vec3) float64 {
	return v.X*u.X + v.Y*u.Y + v.Z*u.Z
}

func Cross(v Vec3, u Vec3) Vec3 {
	return Vec3{
		X: v.Y*u.Z - v.Z*u.Y,
		Y: v.Z*u.X - v.X*u.Z,
		Z: v.X*u.Y - v.Y*u.X,
	}
}

func Unit(v Vec3) Vec3 {
	return DivBy(v, Length(v))
}

func Normalize(v *Vec3) *Vec3 {
	return Div(v, Length(*v))
}
