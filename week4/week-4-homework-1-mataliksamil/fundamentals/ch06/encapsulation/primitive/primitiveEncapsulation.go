package main

import (
	"errors"
	"fmt"
)

// By declaring types for primitives methods can be associated with them.

type PositiveInt int
type String string

//func (i int) whoAreYou() string {
//	return "I am a number nd my value is " + fmt.Sprintf(i)
//}
//
//func (s string) whoAreYou() string {
//	return "I am a number nd my value is " + fmt.Sprintf(i)
//}

func NewI(arg int) (PositiveInt, error) {
	if arg > 0 {
		var i PositiveInt = PositiveInt(arg)
		return i, nil
	}
	return PositiveInt(-1), errors.New("oh you pass an invalid value")
}

// valid validates the value of PositiveInt
func (i PositiveInt) valid() bool {
	if i > 0 {
		return true
	}
	return false
}

func (s String) whoAreYou() string {
	return "I am a string and my value is " + string(s) + ".\n"
}

func main() {
	var i PositiveInt = 5
	fmt.Println(i.valid())
	i = -7
	fmt.Println(i.valid())

	var s String = "Hey, what's up?"
	fmt.Println(s.whoAreYou())
}
