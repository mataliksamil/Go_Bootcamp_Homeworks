package car

import "fmt"

type Car struct {
	make, model string
	year        int
	speed       int
	distance    int
	convertible bool
}

//func New(.....) *Car{
//
//}

func (c *Car) Accelerate(speed int) {
	c.speed = speed
}

func (c *Car) Stop() {
	c.speed = 0
}

func (c *Car) Move(distance int) {
	c.distance += distance
}

func (c *Car) Print() {
	fmt.Printf("%#v\n", c)
}

func main() {
	myCar := Car{"Mercedes", "G", 2022, 0, 0, false}
	myCar.Print()

	myCar.Accelerate(100)
	myCar.Move(200)
	myCar.Print()

	myCar.Accelerate(150)
	myCar.Move(100)
	myCar.Print()

	myCar.Stop()
	myCar.Print()

	fmt.Println()

	yourCar := Car{"BMW", "i8", 2020, 200, 26200, true}
	yourCar.Print()
}
