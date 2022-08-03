package main

import "fmt"

func main() {
	defer panic1()
}

func panic1() int {
	var array [3]string
	fmt.Println(array[0])
	//fmt.Println(array[3])
	i := 3
	fmt.Println(array[i])
	return 4
}
