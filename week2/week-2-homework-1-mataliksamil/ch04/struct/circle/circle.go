package main

import "fmt"

// typical definiton of a struct named point
type Point struct {
	x, y     float64
	distance float64
}

// point struct embedding in Circle1 struct
type Circle1 struct {
	p Point // explicitly
	r int
}

// point struct embedding in Circle2 struct
type Circle2 struct {
	Point // this is an anonymous field
	r     int
}

func main() {
	p1 := Point{x: 3, y: 4}                                  // struct literals
	c1 := Circle1{p1, 10}                                    // p1 assignment can be done directly
	fmt.Printf("c1: %#v\n", c1)                              // #v expression for print the struct with its field can also be printed
	fmt.Printf("c1.p: %#v\n", c1.p)                          // distance is default as zero value since no assignment above
	fmt.Printf("c1.p.x: %#v c1.p.y: %#v \n", c1.p.x, c1.p.y) // we can reach the fields of embedded struct
	//fmt.Printf("c1.p: %#v", c1.Point) // Can't say that! // because we explicitly declared the name of Point struct field as p

	p2 := Point{x: 3, y: 4}
	c2 := Circle2{p2, 10}
	fmt.Printf("c2: %#v\n", c2)
	//fmt.Printf("c2.p: %#v", c2.p) // Can't say that! because the is no explicit declaration as p
	fmt.Printf("c2.p: %#v", c2.Point) // we can reach sub fields by directly calling the name of the struct itself
	fmt.Printf("c2.Point.x: %#v c2.Point.y: %#v \n", c2.Point.x, c2.Point.y)
}
