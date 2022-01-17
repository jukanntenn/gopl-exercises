// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jukanntenn/gopl-exercises/ch4/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s (%s)\n",
			item.Number, item.User.Login, item.Title, age(item.CreatedAt))
	}
}

func age(t time.Time) string {
	dAgo := daysAgo(t)
	if dAgo <= 30 {
		return "less than a month old"
	}
	if dAgo > 30 && dAgo <= 365 {
		return "less than a year old"
	}
	return "more than a year old"
}

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}
