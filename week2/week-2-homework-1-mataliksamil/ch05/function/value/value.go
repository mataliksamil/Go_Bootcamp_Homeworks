package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	greetCreator := createGreetInTurkish
	greetPrinter(greetCreator, "Zeynep")

	greetPrinter(createGreetInTurkish, "Hatice")
	fmt.Println(reflect.TypeOf(greetCreator))

	var greetInEnglish = createGreetInEnglish
	fmt.Println(reflect.TypeOf(greetInEnglish))

	greetPrinter(greetInEnglish, "Mary")
	greetPrinter(convertToUppercase, "naber?")

	var anotherGreetCreator func(string) string = createGreetInTurkish
	greetPrinter(anotherGreetCreator, "Selin")

	//if(createGreetInTurkish == convertToUppercase)
}

func createGreetInTurkish(name string) string {
	return "Selam " + name + " :)"
}

func createGreetInEnglish(name string) string {
	return "Hi " + name + " :)"
}

func convertToUppercase(arg string) string {
	return strings.ToUpper(arg)
}

func greetPrinter(function func(string) string, name string) {
	var greeting = function(name)
	fmt.Printf("%s\n", greeting)
}
