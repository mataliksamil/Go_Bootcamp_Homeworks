package main

import (
	"fmt"
	"time"
)

// the temperature always should be above 100 celcius
var temperature int = 100

func heatUp(i int, t int) {
	fmt.Printf("%d: Temperature before heating up: %d \n", i, temperature)
	time.Sleep(1 * time.Millisecond)
	temperature = temperature + t
	fmt.Printf("%d: Temperature after heating up: %d \n", i, temperature)
}

func coolDown(i int, t int) {
	fmt.Printf("%d: Temperature before the cool down: %d\n", i, temperature)
	// second go routine can get in this else condition since it runs after the check this is race condition
	if temperature-i < 100 {
		fmt.Printf("")
	} else {
		// slowing down the decrease process so to decrease the affect of the race condition
		time.Sleep(1 * time.Millisecond)
		temperature = temperature - t
		fmt.Printf("%d: Temperature after the cooldown: %d\n", i, temperature)

	}
	fmt.Printf("%d: Temperature after heating up: %d \n", i, temperature)
}

func main() {
	fmt.Printf("Temperature at the beginning: %d\n", temperature)
	for i := 1; i <= 10; i++ { // this calls go routines
		go coolDown(i, 20)
		go heatUp(i, 10)
	}
	time.Sleep(1 * time.Second)
	fmt.Printf("Temperature at the end: %d\n", temperature)
}
