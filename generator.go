package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	baseUrl := os.Getenv("JIRA_BASE_URL")
	userName := os.Getenv("JIRA_USER_NAME")
	apiToken := os.Getenv("JIRA_API_TOKEN")
	if len(baseUrl) == 0 || len(userName) == 0 || len(apiToken) == 0 {
		fmt.Println("Required environment variables: $JIRA_BASE_URL, $JIRA_USER_NAME and $JIRA_API_TOKEN.")
		return
	}

	if len(os.Args) < 2 {
		fmt.Println("Required command argument: start commit or tag")
		return
	}
	startCommit := os.Args[1]
	endCommit := ""
	if len(os.Args) > 2 {
		endCommit = os.Args[2]
	}

	commitMessages, err := getCommitMessages(startCommit, endCommit)
	if err != nil {
		log.Fatal(err)
	}
	printArray(Green+"commit messages", commitMessages)

	ticketNumbers, commitsWithoutTicketNumber, err := extractJiraTicketNumbers(commitMessages)
	if err != nil {
		log.Fatal(err)
	}
	if len(commitsWithoutTicketNumber) > 0 {
		printArray(Red+"commit messages without ticket number", commitsWithoutTicketNumber)
	}

	printArray(Green+"jira issues", ticketNumbers)

	details, err := getJiraIssues(ticketNumbers, baseUrl, userName, apiToken)
	if err != nil {
		log.Fatal(err)
	}
	printArray(Green+"issue details", details)
}

const (
	Red   = "\033[31m"
	Green = "\033[32m"
	Reset = "\033[0m"
)

func printArray(title string, data []string) {
	fmt.Println(title + Reset)
	for _, str := range data {
		fmt.Printf("%s\n", str)
	}
	fmt.Println("")
}
