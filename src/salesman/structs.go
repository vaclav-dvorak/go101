package main

import (
	"math"
)

type cityT struct {
	name string
	x, y int
}

func (c *cityT) distance(t cityT) float64 {
	xDis := (c.x - t.x)
	yDis := (c.y - t.y)
	return math.Sqrt(float64((xDis * xDis) + (yDis * yDis)))
}

type routeT struct {
	route    []cityT
	distance float64
}

func (f *routeT) routeDistance() float64 {
	if f.distance != 0 {
		return f.distance
	}
	dist := 0.0
	for k, city := range f.route {
		from := city
		to := f.route[0]
		if k+1 < len(f.route) {
			to = f.route[k+1]
		}
		dist += from.distance(to)
	}
	f.distance = dist
	return f.distance
}

func (f *routeT) toString() string {
	r := ""
	for _, city := range f.route {
		r += city.name
	}
	return r
}
