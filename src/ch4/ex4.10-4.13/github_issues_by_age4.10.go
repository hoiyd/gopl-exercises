package main

import (
	"./github"
	"fmt"
	"log"
	"os"
	"time"
)

const (
	DAY   = 24
	WEEK  = DAY * 7
	MONTH = DAY * 30
	YEAR  = MONTH * 12
)

type categories struct {
	new []*github.Issue
	med []*github.Issue
	old []*github.Issue
}

func categorize(issues *github.IssueSearchResult) *categories {
	category := new(categories)
	for _, issue := range issues.Items {
		t := time.Since(issue.CreatedAt).Hours()
		switch {
		case t < MONTH:
			category.new = append(category.new, issue)
		case t < YEAR:
			category.med = append(category.med, issue)
		default:
			category.old = append(category.old, issue)
		}
	}
	return category
}

func printIssues(issues []*github.Issue) {
	fmt.Printf("%d issues:\n", len(issues))
	for _, item := range issues {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	categorizedResults := categorize(result)
	fmt.Println("Issues created within a month:")
	printIssues(categorizedResults.new)
	fmt.Println("\nIssues created within a year:")
	printIssues(categorizedResults.med)
	fmt.Println("\nIssues created a year before:")
	printIssues(categorizedResults.old)
}
