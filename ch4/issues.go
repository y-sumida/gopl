package main

import (
	"fmt"
	"log"
	"os"
	"gopl/ch4/github"
	"time"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)

	now := time.Now()
	monthAgo := now.AddDate(0, -1, 0)
	yearAgo := now.AddDate(-1, 0, 0)

	fmt.Printf("\n%s 以後\n", monthAgo.Format("2006-01-02"))
	for _, item := range result.Items {
		if item.CreatedAt.After(monthAgo) {
			fmt.Printf("#%-5d %9.9s %.55s %s\n", item.Number, item.User.Login, item.Title, item.CreatedAt)
		}
	}

	fmt.Printf("\n%s 以後\n", yearAgo.Format("2006-01-02"))
	for _, item := range result.Items {
		if item.CreatedAt.After(yearAgo) {
			fmt.Printf("#%-5d %9.9s %.55s %s\n", item.Number, item.User.Login, item.Title, item.CreatedAt)
		}
	}

	fmt.Printf("\n%s 以前\n", yearAgo.Format("2006-01-02"))
	for _, item := range result.Items {
		if item.CreatedAt.Before(yearAgo) {
			fmt.Printf("#%-5d %9.9s %.55s %s\n", item.Number, item.User.Login, item.Title, item.CreatedAt)
		}
	}
}
