package main

import (
	"context"
	"github.com/google/go-github/v47/github"
	"log"
	"sort"
	"strings"
)

func main() {
	owner := "cirruslabs"
	repo := "tart"

	client := github.NewClient(nil)
	issuesOptions := github.IssueListByRepoOptions{}
	issuesOptions.PerPage = 200
	openedIssues, _, err := client.Issues.ListByRepo(context.Background(), owner, repo, nil)
	if err != nil {
		log.Fatal(err)
	}
	issuesOptions.State = "closed"
	closedIssues, _, err := client.Issues.ListByRepo(context.Background(), owner, repo, &issuesOptions)
	if err != nil {
		log.Fatal(err)
	}
	issues := append(openedIssues, closedIssues...)
	log.Println("Found issues", len(issues))

	users := make(map[string]bool)
	for _, issue := range issues {
		users[issue.GetUser().GetLogin()] = true
	}

	commentsOptions := github.IssueListCommentsOptions{
		ListOptions: github.ListOptions{PerPage: 10},
	}
	for {
		comments, commentsResponse, err := client.Issues.ListComments(context.Background(), owner, repo, 0, &commentsOptions)

		if err != nil {
			log.Fatal(err)
		}

		for _, comment := range comments {
			users[comment.GetUser().GetLogin()] = true
		}

		if commentsResponse.NextPage == 0 {
			break
		}
		commentsOptions.Page = commentsResponse.NextPage
	}

	uniqueNames := make([]string, 0, len(users))
	for name := range users {
		uniqueNames = append(uniqueNames, name)
	}
	sort.Slice(uniqueNames, func(i, j int) bool {
		return strings.ToLower(uniqueNames[i]) < strings.ToLower(uniqueNames[j])
	})

	var sb strings.Builder
	for _, username := range uniqueNames {
		sb.WriteString("@")
		sb.WriteString(username)
		sb.WriteString(" ")
	}
	log.Println(sb.String())
}
