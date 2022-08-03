package main

import (
	"fmt"
	"log"
	"os"

	"bootcamp_homeworks/week2/week-2-homework-1-mataliksamil/ch04/extra_json2/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}
