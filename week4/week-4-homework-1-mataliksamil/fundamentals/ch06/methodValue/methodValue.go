package main

import (
	"fmt"
	"time"
)

type Car struct {
	make, model string
	year        int
	speed       int
	distance    int
	convertible bool
}

func (c *Car) start() {
	fmt.Printf("%#v has started!\n", c)
}

func (c *Car) stop() {
	fmt.Printf("%#v has stopped!\n", c)
}

func (c *Car) accelerate(speed int) {
	c.speed = speed
}

func (c *Car) move(distance int) {
	c.distance += distance
}

func main() {

	myCar := &Car{"Mercedes", "G", 2022, 0, 0, false}

	seconds1 := time.Duration(1) * time.Second
	seconds2 := time.Duration(2) * time.Second
	seconds3 := time.Duration(3) * time.Second
	seconds4 := time.Duration(4) * time.Second
	seconds5 := time.Duration(5) * time.Second
	seconds6 := time.Duration(6) * time.Second
	seconds7 := time.Duration(7) * time.Second
	seconds8 := time.Duration(8) * time.Second

	time.AfterFunc(seconds1, func() { myCar.start() })
	time.AfterFunc(seconds2, myCar.stop)

	startMyCar := myCar.start // Method expression
	stopMyCar := myCar.stop   // Method expression
	time.AfterFunc(seconds3, startMyCar)
	time.AfterFunc(seconds4, stopMyCar)

	yourCar := &Car{"BMW", "i8", 2020, 100, 5000, true}

	fStart := (*Car).start                                        // Method expression
	fStop := (*Car).stop                                          // Method expression
	var fAccelerate func(car *Car, speed int) = (*Car).accelerate // Method expression
	var fMove func(car *Car, distance int) = (*Car).move          // Method expression

	fStart(yourCar)
	fStop(yourCar)
	fAccelerate(yourCar, 120)
	fMove(yourCar, 200)

	time.AfterFunc(seconds5, func() { fmt.Printf("Your car: %v\n", yourCar) })

	time.AfterFunc(seconds6, func() { fStart(yourCar) })
	time.AfterFunc(seconds7, func() { fStop(yourCar) })
	time.AfterFunc(seconds8, func() { fmt.Printf("Your car: %v\n", yourCar) })

	time.Sleep(10 * time.Second)
}
