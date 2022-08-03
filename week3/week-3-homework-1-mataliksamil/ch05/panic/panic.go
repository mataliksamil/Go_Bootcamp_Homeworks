package main

import "fmt"

func main() {
	//panic0()
	//panic1()
	//panic2()
	//panic3()
	//panic4()
	// panic5()

}

func panic0() {
	//fmt.Println(0 / 0) // Not a panic, compile time error!
	i := 0
	j := 0
	fmt.Println(i / j)
}

func panic1() {
	var array [3]string
	fmt.Println(array[0])
	//fmt.Println(array[3])
	i := 3
	fmt.Println(array[i])
}

func panic2() {
	p := Person{"Sema"}
	greet(p)
	s := &p
	s = nil
	greet(*s)
}

func panic3() {
	var slice []int = nil
	fmt.Println(slice[0])
}

func panic4() {
	causePanic()
}

func causePanic() {
	var slice []int = nil
	if slice == nil {
		panic("Error: It must not be nil!")
	}
}

type Person struct {
	name string
}

func greet(p Person) {
	fmt.Println("Hi ", p.name)
}

func panic5() {
	var x func()
	x()
}
