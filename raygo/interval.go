package main

import (
	"math"
	"math/rand"
)

type Interval struct {
	Min float64
	Max float64
}

var Universe = Interval{math.Inf(-1), math.Inf(1)}
var Antiverse = Interval{0.0, 0.0}

func (i Interval) size() float64 {
	return i.Max - i.Min
}

func (i Interval) contains(x float64) bool {
	return i.Min <= x && x <= i.Max
}

func (i Interval) surrounds(x float64) bool {
	return i.Min < x && x < i.Max
}

func (i Interval) rand() float64 {
	return i.Min + i.size()*rand.Float64()
}

func (i Interval) clamp(x float64) float64 {
	if x < i.Min {
		return i.Min
	}
	if x > i.Max {
		return i.Max
	}
	return x
}
