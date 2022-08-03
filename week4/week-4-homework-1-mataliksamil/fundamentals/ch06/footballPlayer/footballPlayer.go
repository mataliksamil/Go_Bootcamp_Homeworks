package main

import "fmt"

type FootballPlayer struct {
	name      string
	team      string
	goalCount int
}

func main() {
	salah := FootballPlayer{"Salah", "Liverpool", 18}
	newGoalCount := score(salah)
	salah.goalCount = newGoalCount
	fmt.Printf("%#v\n", salah)
	salah.score1()
	fmt.Printf("%#v\n", salah)
	salah.score2()
	fmt.Printf("%#v\n", salah)
}

func score(fp FootballPlayer) int {
	fp.goalCount++
	return fp.goalCount
}

// this wont increase the goal count of the football player
// since it will copy the football player and increase the goalcount of the copy
// in its own scope so this wont affect global fp "salah"
func (fp FootballPlayer) score1() {
	fp.goalCount++
}

// this will increase the goalcount since it takes football players pointer as input
func (fp *FootballPlayer) score2() {
	fp.goalCount++
}
