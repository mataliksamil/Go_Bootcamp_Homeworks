package main

import (
	another "/../anotherFootballPlayer"
	"fmt"
)

type FootballPlayer struct {
	name      string
	team      string
	goalCount int
}

var akin1 FootballPlayer = FootballPlayer{name: "Akin Kaldiroglu", team: "Ayvalik Gücü", goalCount: 35}

func main() {
	//structs()
	//passToMethods()
	//compare()
	//anotherStruct()
	yetAnotherStruct()
}

func structs() {
	var salah FootballPlayer
	fmt.Println(salah)

	salah.name = "Mo Salah"
	salah.team = "Liverpool"
	salah.goalCount = 18
	fmt.Println(salah)
	fmt.Printf("Name: %s\n", salah.name)
	fmt.Printf("Name: %#v\n", salah)

	debruyne := FootballPlayer{"Kevin Debruyne", "Manchester City", 7}
	fmt.Println(debruyne)

	pp := &FootballPlayer{"Robert Lewandowski", "Bayern Munchen", 28}
	fmt.Println(*pp)
}

func passToMethods() {
	mbappe := FootballPlayer{name: "Kylian Mbappe", team: "PSG", goalCount: 12}
	fmt.Println(mbappe)
	fmt.Printf("Address of MBappe: %p\n", &mbappe)

	score1(mbappe)
	fmt.Printf("MBappe's goal count: %d\n", mbappe.goalCount)

	score2(&mbappe)
	fmt.Printf("MBappe's goal count: %d\n\n", mbappe.goalCount)

}

func score1(fp FootballPlayer) {
	fp.goalCount++
	fmt.Printf("Address of the player: %p\n", &fp)
}

func score2(fp *FootballPlayer) {
	fp.goalCount++
	fmt.Printf("Address of the player: %p\n", fp)
}

func compare() {
	lewo := FootballPlayer{"Robert Lewandowski", "Bayern Munchen", 28}
	anotherLewo := &FootballPlayer{"Robert Lewandowski", "Bayern Munchen", 28}
	yetAnotherLewo := FootballPlayer{"Robert Lewandowski", "Bayern Munchen", 27}

	equal(lewo, *anotherLewo)
	equal(lewo, yetAnotherLewo)
}

// Defines a struct local to a function
func anotherStruct() {
	type FootballPlayer struct {
		name      string
		team      string
		goalCount int
	}

	//var akin2 FootballPlayer = FootballPlayer{team: "Ayvalik Gücü", name: "Akin Kaldiroglu", goalCount: 35}
	//equal(akin1, akin2)
}

// Uses another struct defined another function FootballPlayer
func yetAnotherStruct() {
	fp := another.FootballPlayer{"AyvalikGucu", "Akin", 21}
	fmt.Println(fp)
	fmt.Printf("%T\n", fp)
	//fmt.Println(reflect.TypeOf(fp))
	fmt.Printf("%T\n", akin1)

	//if (fp == akin) {
	//	fmt.Println("The same!")
	//}
}

func equal(fp1 FootballPlayer, fp2 FootballPlayer) {
	if fp1 == fp2 {
		fmt.Println("They are equal")
	} else {
		fmt.Println("They are not equal")
	}
}
