package main

import "fmt"

var i int = 5

func main() {
	//fmt.Println(i)
	//defer deferred1()
	//fmt.Println(i)
	//panic("Ufff!")
	//panic("Ohh, again!")

	//j := 3
	//defer deferred2(j)
	//j += 10
	//fmt.Println(j)

	//k := 3
	//fmt.Println(k)
	//l := defer deferred3(k)
	//fmt.Println(l)

	defer deferred1()
	defer deferred2(1)
	defer deferred3(2)
}

func deferred1() {
	i++
	fmt.Println("In deferred1: ", i)
}

func deferred2(j int) {
	j++
	fmt.Println("In deferred2: ", j)
}

func deferred3(k int) int {
	k++
	fmt.Println("In deferred3: ", k)
	return k
}
