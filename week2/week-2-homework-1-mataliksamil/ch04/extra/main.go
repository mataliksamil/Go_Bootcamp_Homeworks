package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int
	Color  bool
	Actors []string
}
type Bird struct {
	Species     string
	Description string
}

func main() {
	var movies = []Movie{
		{Title: "Casablanca", Year: 1942, Color: false,
			Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Cool Hand Luke", Year: 1967, Color: true,
			Actors: []string{"Paul Newman"}},
		{Title: "Bullitt", Year: 1968, Color: true,
			Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
	}
	fmt.Println("Selamlar")

	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("JSON Marshaling failed: %s", err)
	}
	err = nil
	fmt.Printf("%s", data)
	fmt.Println("___________________________________")
	fmt.Println("------------------------------------")

	//	HOW TO UNMARSHALL AFTER MARSHALINDENT FUNCTION
	var titles []struct{ Title string }
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Println(titles) // "[{Casablanca} {Cool Hand Luke} {Bullitt}]"

}
