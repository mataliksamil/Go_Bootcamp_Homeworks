package main

import (
	"fmt"
	mathFunction "fundamentals/ch07/math"
)

// That's about dynamic dispatching when calling methods on an interface type in Go
func main() {

	var mf mathFunction.MathFunction = mathFunction.MathFunctionFactory()
	val := mf.Calculate(0.784) // That's pi/4
	fmt.Printf("%f\n", val)    // Which method is called? sin, cos or log
	//fmt.Printf("%T", mf)
}
