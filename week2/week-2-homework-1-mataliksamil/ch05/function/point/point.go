package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y int
}

func main() {
	p1 := Point{3, 4}
	xAxis1, yAxis1, distanceToOrigin1 := pointProperties(p1)
	fmt.Println(distanceToOrigin1, xAxis1, yAxis1)

	p2 := Point{-12, -13}
	_, _, distanceToOrigin2 := pointProperties(p2)
	fmt.Println(distanceToOrigin2)
}

func pointProperties(p Point) (int, int, float64) {
	//var distance float64 = math.Sqrt(float64(p.x*p.x + p.y*p.y))
	var distance float64 = math.Sqrt(math.Pow(float64(p.x), 2) + math.Pow(float64(p.y), 2))
	return p.x, p.y, distance
}
