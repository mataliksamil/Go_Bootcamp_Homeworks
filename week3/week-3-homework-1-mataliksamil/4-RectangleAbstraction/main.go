package main

import (
	"errors"
	"fmt"
	"os"
)

type Rectangle struct {
	Height float64
	Length float64
}

// returns the area of given rectangle as float64
func area(myRectangle Rectangle) float64 {
	return myRectangle.Length * myRectangle.Height
}

// returns the circumference of given rectangle as float64
func circumference(myRectangle Rectangle) float64 {
	return 2 * (myRectangle.Length + myRectangle.Height)
}

// creates a rectangle type object. throws an error for the inconvenient input parameters
func getRectangle(myHeight float64, myLength float64) (Rectangle, error) {
	var myRectangle Rectangle // default object
	if myHeight <= 0 {
		fmt.Println(myHeight, "---------------")
		return myRectangle, errors.New("ERROR! Negative Height \nHeight :" + fmt.Sprintf("%.1f", myHeight)) // throws error
	}
	if myLength <= 0 {

		return myRectangle, errors.New("ERROR! Negative Length \nLength :" + fmt.Sprintf("%.1f", myLength)) // throws error
	}
	myRectangle = Rectangle{myHeight, myLength} // creation of rectangle object
	return myRectangle, nil                     // nil error returns since there is no error
}

func main() {

	newRectangle, err := getRectangle(-1, 16) // function call
	if err != nil {                           // catches the error in the main function so it can show the error message
		fmt.Println(err)
		os.Exit(1) // in order to terminate the program after the error
	}

	fmt.Println("Area of new rectangle : ", area(newRectangle))
	fmt.Println("Circumference of new rectangle : ", circumference(newRectangle))

}
