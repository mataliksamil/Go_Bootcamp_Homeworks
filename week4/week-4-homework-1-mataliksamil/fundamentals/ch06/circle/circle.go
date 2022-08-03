package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y, distance float64
}

func (p *Point) calculateDistanceToOrigin() float64 {
	p.distance = math.Sqrt(p.x*p.x + p.y*p.y)
	return p.distance
}

func (p *Point) calculateDistanceTo(q Point) float64 {
	distance := math.Sqrt(math.Pow(p.x-q.x, 2) + math.Pow(p.y-q.y, 2))
	return distance
}

func (p *Point) move(direction rune, amount float64) {
	if direction == 'x' {
		p.x += amount
	} else if direction == 'y' {
		p.y += amount
	}
	p.calculateDistanceToOrigin()
}

func (p *Point) print() {
	fmt.Printf("%#v\n", p)
}

func main() {
	processPoint()
	//processCircle()

}

func processPoint() {
	p1 := new(Point) // Returns a pointer to Point instance
	p1.print()

	p1.move('x', 5)
	p1.move('y', 11)
	p1.print()

	p1.move('x', 4)
	p1.move('y', 1)
	p1.print()

	p2 := Point{x: 3, y: 5}
	fmt.Printf("Distance is %f\n", p1.calculateDistanceTo(p2))
	fmt.Printf("Distance is %f\n", p2.calculateDistanceTo(*p1)) // since p1 called as pointer
}

func processCircle() {
	p1 := Point{x: 3, y: 4}
	p1.calculateDistanceToOrigin()
	fmt.Printf("Distance of %#v to the origin is %f\n", p1, p1.distance)
	c1 := Circle{p1, 10}
	c1.print()
	// It is delegated to Point of Circle.
	fmt.Printf("Distance of %#v to the origin is %f\n", c1, c1.calculateDistanceToOrigin())

	c1.move('x', -3)
	c1.print()
	c1.move('y', -4)
	c1.print()
}

type Circle struct {
	Point
	r int
}

// Following two methods are the same.
func (c *Circle) calculateDistance1() float64 {
	c.distance = math.Sqrt(c.x*c.x + c.y*c.y)
	return c.distance
}

func (c *Circle) calculateDistance2() float64 {
	c.distance = math.Sqrt(c.Point.x*c.Point.x + c.Point.y*c.Point.y)
	return c.distance
}

// Following two methods are the same.
func (c *Circle) calculateDistance3() float64 {
	return c.calculateDistanceToOrigin()
}

func (c *Circle) calculateDistance4() float64 {
	return c.Point.calculateDistanceToOrigin()
}

// Comment out this method and see what happens. When this method commented out move() of Point is called.
func (c *Circle) move(direction rune, amount float64) {
	c.Point.move(direction, amount)
}

// Comment out this method and see what happens. When this method commented out print() of Point is called.
func (c *Circle) print() {
	fmt.Printf("%#v\n", c)
}
