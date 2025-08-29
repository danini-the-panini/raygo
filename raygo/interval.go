package main

import "math"

type Interval struct {
	Min float64
	Max float64
}

var Universe = Interval{math.Inf(-1), math.Inf(1)}
var Antiverse = Interval{0.0, 0.0}

func (self Interval) size() float64 {
	return self.Max - self.Min
}

func (self Interval) contains(x float64) bool {
	return self.Min <= x && x <= self.Max
}

func (self Interval) surrounds(x float64) bool {
	return self.Min < x && x < self.Max
}
