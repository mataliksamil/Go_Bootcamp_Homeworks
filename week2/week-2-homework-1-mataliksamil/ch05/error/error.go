package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"os"
)

const LOG_FILE = "/Users/akin/Desktop/GoLog.txt"

func init() {
	logFile, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Panic(err)
	}
	//defer logFile.Close()
	log.SetOutput(logFile)
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

type Circle struct {
	r float64
}

func main() {
	//// strategy1a
	//err1 := strategy1a()
	//if err1 != nil {
	//	fmt.Printf("Error: %s\n", err1)
	//}
	//
	//// strategy1b
	//err2 := strategy1b()
	//if err2 != nil {
	//	fmt.Printf("Error: %s\n", err2)
	//}

	//strategy2()
	//strategy3()
	//strategy4()
	// strategy5()
}

func strategy1a() error {
	fmt.Println("\nstrategy1a")
	c1 := Circle{r: 12}
	circ1, err := circumference(c1)
	if err != nil {
		return err
	} else {
		fmt.Printf("Circumference is %f\n", circ1)
		return nil
	}
}

func strategy1b() error {
	fmt.Println("\nstrategy1b")
	c1 := Circle{r: -12}
	circ1, err := circumference(c1)
	if err != nil {
		return fmt.Errorf("hey, we've a problem, this is what I've got: %v", err)
	} else {
		fmt.Printf("Circumference is %f\n", circ1)
		return nil
	}
}

func strategy2() {
	fmt.Println("\nstrategy2")
	c1 := Circle{r: -12}
	circ1, err := circumference(c1)
	if err != nil {
		c1 = Circle{r: 8}
		circ1, _ = circumference(c1)
		fmt.Printf("Circumference is %f\n", circ1)
	} else {
		fmt.Printf("Circumference is %f\n", circ1)
	}
}

func strategy3() {
	fmt.Println("\nstrategy3")
	c1 := Circle{r: -12}
	circ1, err := circumference(c1)
	if err != nil {
		fmt.Println("I am quitting!")
		os.Exit(0)
	} else {
		fmt.Printf("Circumference is %f\n", circ1)
	}
}

func strategy4() {
	fmt.Println("\nstrategy4")
	c1 := Circle{r: -12}
	circ1, err := circumference(c1)
	if err != nil {
		log.Println("Calculation failed: ", err)
	} else {
		fmt.Printf("Circumference is %f\n", circ1)
	}
}

func strategy5() {
	fmt.Println("\nstrategy5")
	c1 := Circle{r: -12}
	circ1, _ := circumference(c1)
	if circ1 == -1 {
		fmt.Printf("Oh, that's problem!")
	} else {
		fmt.Printf("Circumference is %f\n", circ1)
	}
}

func circumference(c Circle) (float64, error) {
	if c.r < 0 {
		return -1, errors.New("Negative radius: " + fmt.Sprintf("%f", c.r))
	} else {
		circ := math.Pi * math.Pow(c.r, 2)
		return circ, nil
	}
}
