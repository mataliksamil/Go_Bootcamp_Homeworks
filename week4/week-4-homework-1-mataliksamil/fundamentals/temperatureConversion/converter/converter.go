package main

import (
	"errors"
	"fmt"
	temperature "fundamentals/temperatureConversion/temperature"
	"os"
	"strings"
)

var flag bool = true

type Converter struct {
	functions []temperature.Temperature
}

func (c *Converter) addConversion(t temperature.Temperature) {
	c.functions = append(c.functions, t)
}

func (c *Converter) doConversion(name string, arg float64) (float64, error) {
	var result float64
	for _, f := range c.functions {
		// this line converts input operation string as all lower case so inputs case-insensitive
		if strings.ToLower(name) == strings.ToLower(f.GetName()) {
			result = f.ConvertTo(arg)
			return result, nil
		}
	}
	return 0, errors.New("no such conversion exists:" + name)
}
func main() {
	//functions()
	//converter()
	startConverter()
}

func startConverter() {
	myConverter := Converter{}
	myConverter.addConversion(temperature.CelciusToFahrenheit{"celciustofahrenheit"})
	myConverter.addConversion(temperature.CelciusToKelvin{"celciustokelvin"})
	myConverter.addConversion(temperature.FahrenheitToCelcius{"fahrenheittocelcius"})
	myConverter.addConversion(temperature.FahrenheitToKelvin{"fahrenheittokelvin"})
	myConverter.addConversion(temperature.KelvinToCelcius{"kelvintocelcius"})
	myConverter.addConversion(temperature.KelvinToFahrenheit{"kelvintofahrenheit"})

	fmt.Println("\nConverter started.")
	fmt.Println("You can calculate using following functions")
	for _, f := range myConverter.functions { // iterate through defined math functions above
		fmt.Println(f.GetName())
	}

	for flag { // loop until flag down
		flag = true
		var type1Name string
		var type2Name string
		var temp1 float64

		fmt.Println("> Enter <type> <temperature> <convert this type> with blank between them")
		_, err := fmt.Scanf("%s", &type1Name)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		//
		if type1Name == "x" {
			//os.Exit(0)
			flag = false
		} else {
			fmt.Println(">")
			_, err := fmt.Scanf("%f", &temp1)
			if err != nil {
				fmt.Println(err)
				os.Exit(0)
			}

			// this is for the type conversion so we can convert
			fmt.Println(">")
			j, err := fmt.Scanf("%s", &type2Name)
			if err != nil {
				fmt.Println(err)
				os.Exit(0)
			}
			j++
			// this is the way we check the string for each possible conversion
			cName := type1Name + "to" + type2Name

			value, err := myConverter.doConversion(cName, temp1)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Result of %s (%v) to %s is %v\n", type1Name, temp1, type2Name, value)
				fmt.Println()
				fmt.Println()
			}

		}
	}
	println("BYE !")

}
