package main

import (
	"fmt"
	"os"
)

type JiraIssueInfo struct {
	Key    string `json:"key"`
	Fields struct {
		Summary string `json:"summary"`
	} `json:"fields"`
}

func pullDetails(startCommit, endCommit, baseURL, userName, apiToken string) ([]string, error) {
	fmt.Printf("getting commit messages from %s to %s\n", startCommit, endCommit)
	commitMessages, err := getCommitMessages(startCommit, endCommit)
	if err != nil {
		return nil, err
	}
	fmt.Printf("commit messages: %+q\n", commitMessages)

	fmt.Printf("extracting JIRA issue numbers from commit messages\n")
	ticketNumbers, err := extractJiraTicketNumbers(commitMessages)
	if err != nil {
		return nil, err
	}
	fmt.Printf("issue numbers: %+q\n", ticketNumbers)

	fmt.Printf("pulling JIRA issue details\n")
	changes, err := getJiraIssues(ticketNumbers, baseURL, userName, apiToken)
	if err != nil {
		return nil, err
	}

	return changes, nil
}

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

	changes, _ := pullDetails(startCommit, endCommit, baseUrl, userName, apiToken)
	for _, change := range changes {
		fmt.Println(change)
	}

}
